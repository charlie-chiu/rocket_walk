let type = "WebGL"
if (!PIXI.utils.isWebGLSupported()) {
    type = "canvas"
}
PIXI.utils.sayHello(type)

let app = new PIXI.Application({
    width: 600,
    height: 800,
    antialias: true,
})
app.renderer.backgroundColor = 0x061639

let sky = newSky()
app.stage.addChild(sky);
let rocket = newRocket()
app.stage.addChild(rocket);

let scrollSky = (delta) => {
    sky.tilePosition.x -=  10 * delta;
    sky.tilePosition.y += 20 * delta;
};
app.ticker.add(scrollSky);

function newRocket() {
    let rocketTexture = PIXI.Texture.from('images/rocket.png');
    let rocket = new PIXI.Sprite(
        rocketTexture,
    )
    rocket.transform.scale.set(0.5, 0.5)
    rocket.rotation = 0.4
    rocket.anchor.set(0.5, 1)
    rocket.x = app.screen.width/2 - 100
    rocket.y = app.screen.height  - 100
    return rocket
}

function newSky() {
    let skyTexture = PIXI.Texture.from('images/sky.png');
    return new PIXI.TilingSprite(
        skyTexture,
        app.screen.width,
        app.screen.height,
    )
}

window.onload = function () {
    let conn;
    if (window["WebSocket"]) {
        conn = new WebSocket("ws://" + document.location.host + "/rocketrun");
        conn.onmessage = function (evt) {
            rocket.x += 5
            rocket.y -= 5
            // var messages = evt.data.split('\n');
        };
    } else {
        const item = document.createElement("div");
        item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
    }
};