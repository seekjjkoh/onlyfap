interface Observer<Input> {
	next(value: Input): void;
	complete(): void;
	unsub?: () => void;
}
class SafeObserver<T> implements Observer<T> {
	// constructor enforces that we are always subscribed to destination
	private isUnsubscribed = false;
	private destination: Observer<T>;

	constructor(destination: Observer<T>) {
		this.destination = destination;
		if (destination.unsub) {
			this.unsub = destination.unsub;
		}
	}
	next(value: T): void {
		if (!this.isUnsubscribed) {
			this.destination.next(value);
		}
	}
	complete(): void {
		if (!this.isUnsubscribed) {
			this.destination.complete();
			this.unsubscribe();
		}
	}
	unsubscribe(): void {
		if (!this.isUnsubscribed) {
			this.isUnsubscribed = true;
			if (this.unsub) this.unsub();
		}
	}
	unsub?: () => void;
}

class Observable<Input> {
	constructor(private _subscribe: (_: Observer<Input>) => () => void) { }
	subscribe(next: (_: Input) => void, complete?: () => void): () => void {
		const safeObserver = new SafeObserver(<Observer<Input>>{
			next: next,
			complete: complete ? complete : () => console.log('complete')
		});
		safeObserver.unsub = this._subscribe(safeObserver);
		return safeObserver.unsubscribe.bind(safeObserver);
	}
	static fromEvent<E extends Event>(el: Node, name: string): Observable<E> {
		return new Observable<E>((observer: Observer<E>) => {
			const listener = <EventListener>((e: E) => observer.next(e));
			el.addEventListener(name, listener);
			return () => el.removeEventListener(name, listener);
		})
	}
	static fromArray<V>(arr: V[]): Observable<V> {
		return new Observable<V>((observer: Observer<V>) => {
			arr.forEach(el => observer.next(el));
			observer.complete();
			return () => { };
		});
	}
	static interval(milliseconds: number): Observable<number> {
		return new Observable<number>(observer => {
			let elapsed = 0;
			const handle = setInterval(() => observer.next(elapsed += milliseconds), milliseconds)
			return () => clearInterval(handle);
		})
	}
	map<R>(transform: (_: Input) => R): Observable<R> {
		return new Observable<R>(observer =>
			this.subscribe(e => observer.next(transform(e)), () => observer.complete())
		);
	}
	forEach(f: (_: Input) => void): Observable<Input> {
		return new Observable<Input>(observer =>
			this.subscribe(e => {
				f(e);
				return observer.next(e);
			},
				() => observer.complete()))
	}
	filter(condition: (_: Input) => boolean): Observable<Input> {
		// Your code here ...
		return new Observable<Input>(observer =>
			this.subscribe(e => {
				if (condition(e)) observer.next(e)
			},
				() => observer.complete()));
	}
	takeUntil<V>(o: Observable<V>): Observable<Input> {
		return new Observable<Input>(observer => {
			const oUnsub = o.subscribe(_ => {
				observer.complete();
				oUnsub();
			});
			return this.subscribe(e => observer.next(e), () => {
				observer.complete();
				oUnsub();
			});
		});
	}
	flatMap<Output>(streamCreator: (_: Input) => Observable<Output>): Observable<Output> {
		return new Observable<Output>((observer: Observer<Output>) => {
			return this.subscribe(t => streamCreator(t).subscribe(o => observer.next(o)), () => observer.complete())
		})
	}
	scan<V>(initialVal: V, fun: (a: V, el: Input) => V): Observable<V> {
		return new Observable<V>((observer: Observer<V>) => {
			let accumulator = initialVal;
			return this.subscribe(
				v => {
					accumulator = fun(accumulator, v);
					observer.next(accumulator);
				},
				() => observer.complete()
			)
		})
	}
}
