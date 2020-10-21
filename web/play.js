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

app.ticker.add((delta) => {

    sky.tilePosition.x +=  1 * delta;
    sky.tilePosition.y += 20 * delta;
});

function newRocket() {
    let rocketTexture = PIXI.Texture.from('images/rocket.png');
    let rocket = new PIXI.Sprite(
        rocketTexture,
    )
    rocket.transform.scale.set(0.5, 0.5)
    // rocket.rotation = 0.1
    rocket.anchor.set(0.5, 1)
    rocket.x = app.screen.width/2
    rocket.y = app.screen.height
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