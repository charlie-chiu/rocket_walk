window.onload = function () {
    handleWS()
};


let skyVectorX = 0.05;
let skyVectorY = 0.05;

let app = initPIXI();
let sky = newSky()
app.stage.addChild(sky);
app.ticker.add(scrollSky);
let rocket = newRocket()
app.stage.addChild(rocket);
let explosion = newExplosion()
app.stage.addChild(explosion);

function scrollSky(delta) {
    sky.tilePosition.x -= skyVectorX * delta;
    sky.tilePosition.y += skyVectorY * delta;
}
function increaseSkyVector(limit = 25) {
    while (skyVectorY <= 50) {
        let rand = (Math.random() - 0.5) * 0.1
        skyVectorX += rand + 0.05
        skyVectorY += rand + 0.10
    }
}
function decreaseSkyVector(limit = 25) {
    skyVectorX *= 0.2
    skyVectorY *= 0.2
}
function stopSky() {
    skyVectorX = 0.05
    skyVectorY = 0.05
}

function initPIXI() {
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

    return app
}
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

function handleWS() {
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
function onState(payload) {
    switch (payload.name) {
        case "ready":
            stopSky()
            rocket.visible = true
            break
        case "betend":
            break
        case "launch":
            increaseSkyVector()
            break
        case "bust":
            explosion.visible = true
            rocket.visible = false
            decreaseSkyVector()
            break
        case "end":
            explosion.visible = false
            break
    }
}