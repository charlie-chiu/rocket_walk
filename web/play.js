let type = "WebGL"
if (!PIXI.utils.isWebGLSupported()) {
    type = "canvas"
}
PIXI.utils.sayHello(type)

let application = PIXI.Application;
let loader = PIXI.Loader.shared;

let app = new application({
    width:800,
    height:600,
    antialias: true,
})
app.renderer.backgroundColor = 0x061639
let rocket = new PIXI.Sprite
loader.add("images/rocket.png").load(setup)

function setup() {
    rocket.texture = loader.resources["images/rocket.png"].texture;
    rocket.name = "rocket"
    rocket.transform.scale.set(0.5,0.5)
    rocket.rotation = 0.5
    rocket.anchor.set(0.5,0.5)
    rocket.x = 400
    rocket.y = 300
    app.stage.addChild(rocket)
}

console.log(rocket)

rocket.y = 500
