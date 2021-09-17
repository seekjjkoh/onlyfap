from typing import Dict
from time import sleep
import json

import luigi

def sleepn(time: int):
	sleep(time)


def save_checkpoint(filename: str, content: Dict):
	with open(filename, "w") as f:
		json.dump(content, f)


class Task1(luigi.Task):
	def run(self):
		sleep(10)
		save_checkpoint("task1.json", {})

	def output(self):
		return luigi.LocalTarget("task1.json")


class Task2(luigi.Task):
	def requires(self):
		return Task1()

	def run(self):
		sleep(20)
		save_checkpoint("task2.json", {})
	
	def output(self):
		return luigi.LocalTarget("task2.json")


class Task25(luigi.Task):
	def requires(self):
		return Task1()

	def run(self):
		sleep(25)
		save_checkpoint("task25.json", {})
	
	def output(self):
		return luigi.LocalTarget("task25.json")


class Task3(luigi.Task):
	def requires(self):
		return [Task2(), Task25()]

	def run(self):
		sleep(30)
		save_checkpoint("task3.json", {})
	
	def output(self):
		return luigi.LocalTarget("task3.json")


class Task4(luigi.Task):
	def requires(self):
		return Task3()

	def run(self):
		sleep(40)
		save_checkpoint("task4.json", {})
	
	def output(self):
		return luigi.LocalTarget("task4.json")

if __name__ == "__main__":
	# remember to run luigid
	luigi.build([Task1(), Task2(), Task25(), Task3(), Task4()], local_scheduler=False)
	luigi.run()
