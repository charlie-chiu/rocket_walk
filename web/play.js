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
let explosion = newExplosion()
app.stage.addChild(explosion);

let scrollSky = (delta) => {
    sky.tilePosition.x -= 25 * delta;
    sky.tilePosition.y += 50 * delta;
};

function newRocket() {
    let rocketTexture = PIXI.Texture.from('images/rocket.png');
    let rocket = new PIXI.Sprite(
        rocketTexture,
    )
    rocket.transform.scale.set(0.5, 0.5)
    rocket.rotation = 0.4
    rocket.anchor.set(0.5, 1)
    rocket.x = app.screen.width / 2 - 200
    rocket.y = app.screen.height
    return rocket
}
function newExplosion() {
    let explosionTexture = PIXI.Texture.from('images/explosion.png');
    let explosion = new PIXI.Sprite(
        explosionTexture,
    )
    explosion.alpha = 0.9
    explosion.transform.scale.set(0.3,0.3)
    explosion.rotation =1.5
    explosion.anchor.set(0.5, 0.5)
    explosion.x = app.screen.width / 2 - 150
    explosion.y = app.screen.height - 150
    explosion.visible = false
    return explosion
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
    handleWS()
};

let handleWS = () => {
    let conn;
    if (window["WebSocket"]) {
        conn = new WebSocket("ws://" + document.location.host + "/rocketrun");
        conn.onmessage = function (evt) {
            let data = JSON.parse(evt.data)
            if (data.name === "on_state") {
                onState(data.payload)
            }
        };
    } else {
        const item = document.createElement("div");
        item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
    }
}
let onState = (payload) => {
    switch (payload.name) {
        case "new":
            explosion.visible = false
            console.log("state: new")
            break
        case "betend":
            console.log("state: betend")
            break
        case "launch":
            console.log("state: launch")
            app.ticker.add(scrollSky);
            break
        case "bust":
            explosion.visible = true
            console.log("state: bust");
            break
        case "end":
            console.log("state: end")
            app.ticker.remove(scrollSky);
            break
    }
}