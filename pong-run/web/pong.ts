// pong.ts

// const url = `ws://localhost:8080/ws${window.location.search}`;
const url = `wss://${window.location.host}/ws${window.location.search}`;
let ws: WebSocket;
let player = 0;
let svg: HTMLElement;
let ball: Elem;
let p1: Elem;
let p2: Elem;
let gameState = 0;
let gameScore = {
  p1: 0,
  p2: 0,
};
let lastMouseEvent = Date.now();

function pong() {
  const svg = document.getElementById("canvas")!;
  new Elem(svg, "rect")
    .attr("x", 297).attr("y", 20)
    .attr("width", 6).attr("height", 80)
    .attr("fill", "#95B3D7");
  new Elem(svg, "rect")
    .attr("x", 297).attr("y", 140)
    .attr("width", 6).attr("height", 80)
    .attr("fill", "#95B3D7");
  new Elem(svg, "rect")
    .attr("x", 297).attr("y", 260)
    .attr("width", 6).attr("height", 80)
    .attr("fill", "#95B3D7");
  new Elem(svg, "rect")
    .attr("x", 297).attr("y", 380)
    .attr("width", 6).attr("height", 80)
    .attr("fill", "#95B3D7");
  new Elem(svg, "rect")
    .attr("x", 297).attr("y", 500)
    .attr("width", 6).attr("height", 80)
    .attr("fill", "#95B3D7");
  setup();
  setupWs();
}

function setup() {
  svg = document.getElementById("canvas")!;
  // create a ball
  ball = new Elem(svg, "circle")
    .attr("cx", 60).attr("cy", 95)
    .attr("r", 8)
    .attr("fill", "#95B3D7")
    .attr("id", "ball");
  p1 = new Elem(svg, "rect")
    .attr("x", 50).attr("y", 70)
    .attr("width", 8).attr("height", 100)
    .attr("fill", "#95B3D7");
  p2 = new Elem(svg, "rect")
    .attr("x", 550).attr("y", 70)
    .attr("width", 8).attr("height", 100)
    .attr("fill", "#95B3D7");
};

function setupWs() {
  ws = new WebSocket(url);
  ws.onopen = () => {
    ws.addEventListener("message", wsHandler);
  };
  ws.onclose = () => {
    ws.removeEventListener("message", wsHandler);
  };
}

function wsHandler(event: MessageEvent<any>) {
  const pongData = JSON.parse(event.data);
  if (pongData.player) {
    player = pongData.player;
    registerPlayer(player);
  } if (pongData.action) {
    processAction(pongData);
  }
};

function registerPlayer(player: number) {
  if (player != 1 && player != 2) {
    return;
  }
  Observable
    .fromEvent<MouseEvent>(document, "mousemove")
    .map(({ clientX, clientY }) => ({
      x: clientX,
      y: clientY - 150, // 150px (title)
    }))
    .subscribe(({ x, y }) => {
      if (y > 500) {
        y = 500;
      } else if (y < 0) {
        y = 0;
      }
      // emit to websocket
      const newNow = Date.now()
      if (newNow - lastMouseEvent < 50) {
        return;
      }
      lastMouseEvent = newNow;
      ws.send(JSON.stringify({ action: "move", data: { type: "player", player, x, y } }));
      if (player != 1) {
        return;
      }
      if (gameState == 0) {
        emitBallPosition(Number(p1.attr("x")) + 10, y + 50);
      }
    });
  Observable.fromEvent<MouseEvent>(document, "mousedown")
    .subscribe(() => {
      // emit to websocket
      ws.send(JSON.stringify({ action: "click", data: { type: "player", player } }));
      if (player != 1) {
        return;
      }
      handleStates();
    });
};

function randomSign(): number {
  if (Math.random() > 0.5) {
    return -1
  }
  return 1
}

function emitBallPosition(x: number, y: number) {
  ws.send(JSON.stringify({ action: "move", data: { type: "ball", x, y } }));
}

function handleStates() {
  if (gameState == 0) {
    gameState = 1;
    let xVar = 10;
    let yVar = randomSign() * (5 + Math.random() * 5);
    const animate = window.setInterval(() => {
      if (Number(ball.attr('cx')) == 550
        && (Number(ball.attr('cy')) >= Number(p2.attr('y')) && Number(ball.attr('cy')) <= Number(p2.attr('y')) + 100)) {
        xVar = -10;
        yVar = 5 + Math.ceil(Math.random() * 5);
      } else if (Number(ball.attr('cx')) == 50
        && (Number(ball.attr('cy')) >= Number(p1.attr('y')) && Number(ball.attr('cy')) <= Number(p1.attr('y')) + 100)) {
        xVar = 10;
        yVar = randomSign() * Math.ceil(5 + Math.random() * 5);
      } else if (Number(ball.attr('cy')) <= 8 || Number(ball.attr('cy')) >= 592) {
        // bounce
        yVar = -1 * yVar;
      } else if (Number(ball.attr('cx')) <= 10) {
        gameScore.p2++
        gameState = 0;
        broadcastState("p2 scored");
        clearInterval(animate);
      } else if (Number(ball.attr('cx')) >= 590) {
        gameScore.p1++
        gameState = 0;
        broadcastState("p1 scored");
        clearInterval(animate);
      }
      emitBallPosition(xVar + Number(ball.attr("cx")), yVar + Number(ball.attr("cy")))
    }, 50);
  }
}

function broadcastState(log: string) {
  ws.send(JSON.stringify({action: "score", data: gameScore}));
  ws.send(JSON.stringify({action: "log", data: {log}}));
}

function processAction(wsData: any) {
  // nothing interesting
  if (wsData.action == "click") {
    return;
  }
  const { data } = wsData;
  if (data.x && data.y) {
    if (data.type == "ball") {
      ball.attr("cx", data.x);
      ball.attr("cy", data.y);
      return;
    }
    const ele = data.player == 1 ? p1 : p2;
    ele.attr("y", data.y);
    return;
  }
  if (wsData.action == "score") {
    document.getElementById("score")!.innerHTML = `${data.p1} : ${data.p2}`;
    return;
  }
  if (wsData.action == "log") {
    document.getElementById("gamelog")!.innerHTML = data.log;
    return;
  }
};

if (typeof window != 'undefined')
  window.onload = () => {
    pong();
  }