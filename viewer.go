package main

import "github.com/icza/gowut/gwu"

func NewModelViewer(actor *Actor) gwu.HTML {
 	var Mesh, Texture, TextureLoader string
 	var Normal, Specular, AmbientOcclusion string
	
	if actor == nil || actor.Mesh() == "" {
		return  gwu.NewHTML(``)
	}
	
	var Github = "https://raw.githubusercontent.com/0ad/0ad/master/binaries/data/mods/public/"
		
	
	Mesh = "https://raw.githubusercontent.com/0ad/0ad/master/binaries/data/mods/public/"+actor.Mesh()
	Texture = "https://raw.githubusercontent.com/0ad/0ad/master/binaries/data/mods/public/"+actor.Texture()
	
	if actor.Normal() != "" {
		Normal = `normalMap: textureLoader.load('`+Github+actor.Normal()+`'),`
	}
	if actor.Specular() != "" {
		Specular = `specularMap: textureLoader.load('`+Github+actor.Specular()+`'),`
	}
	if actor.AmbientOcclusion() != "" {
		AmbientOcclusion = `aoMap: textureLoader.load('`+Github+actor.AmbientOcclusion()+`'),`
	}

	var PropLoaderString string
	Props := actor.Props()
	for prop, _ := range Props {
		
		var Mesh, Texture, TextureLoader string
		var Normal, Specular, AmbientOcclusion string
		
		actor, err := LoadActor(prop)
		if err != nil {
			continue
		}

		if len(Texture) > 0 {
			if Texture[len(Texture)-4:] == ".dds" {
				TextureLoader = "THREE.DDSLoader()"
			} else {
				TextureLoader = "THREE.TextureLoader()"	
			}
		} else {
			TextureLoader = "THREE.TextureLoader()"	
		}
		
		if actor.Mesh() == "" {
			continue
		}
		
		Mesh = Github+actor.Mesh()
		Texture = Github+actor.Texture()
		
		if actor.Normal() != "" {
			Normal = `normalMap: textureLoader.load('`+Github+actor.Normal()+`'),`
		}
		if actor.Specular() != "" {
			Specular = `specularMap: textureLoader.load('`+Github+actor.Specular()+`'),`
		}
		if actor.AmbientOcclusion() != "" {
			AmbientOcclusion = `aoMap: textureLoader.load('`+Github+actor.AmbientOcclusion()+`'),`
		}
		
		PropLoaderString += `
		
		if (true) {
			let texture = (new `+TextureLoader+`).load('`+Texture+`');
					
			texture.wrapS = THREE.RepeatWrapping;
			texture.wrapT = THREE.RepeatWrapping;
			
			let textureLoader = new THREE.TextureLoader();
			
			let material = new THREE.MeshPhongMaterial({
				map: texture,
				`+Normal+`
				`+Specular+`
				`+AmbientOcclusion+`
			});
			
			let loader = new THREE.ColladaLoader( loadingManager );
				loader.load( '`+Mesh+`', function ( collada ) {

					collada.scene.traverse(function (node) {
						if (node.isMesh) {
							node.material.map = material.map;
						}
					});

					models.push(collada.scene);
					

			} );
		}`

	}
	
	if len(Texture) > 0 {
		if Texture[len(Texture)-4:] == ".dds" {
			TextureLoader = "THREE.DDSLoader()"
		} else {
			TextureLoader = "THREE.TextureLoader()"	
		}
	} else {
		TextureLoader = "THREE.TextureLoader()"	
	}
	
	return gwu.NewHTML(`
		<div id="container" style="position:fixed;top:calc(100vh - 400px);left:0px;background-color:black; width:400px;height:400px;color:white;"></div>

		<script src="js/three.js"></script>

		<script src="js/loaders/ColladaLoader.js"></script>
		<script src="js/loaders/DDSLoader.js"></script>
		<script src="js/Detector.js"></script>

		<script>

			//if ( ! Detector.webgl ) Detector.addGetWebGLMessage();

			var container, stats, clock;
			var camera, scene, renderer, elf;
			
			var models = [];

			init();
			animate();

			function init() {

				container = document.getElementById( 'container' );

				camera = new THREE.PerspectiveCamera( 45, window.innerWidth / window.innerHeight, 0.1, 2000 );
				camera.position.set( 2, 1.5, 0.25 );
				camera.lookAt( new THREE.Vector3( 0, 0.5, 0 ) );

				scene = new THREE.Scene();

				clock = new THREE.Clock();

				// loading manager

				var loadingManager = new THREE.LoadingManager( function() {
					
					for (var i = 0; i < models.length; i++) {
						models[i].scale.set(0.1, 0.1, 0.1);
						scene.add( models[i] );
					}

				} );

				// collada

				var texture = (new `+TextureLoader+`).load('`+Texture+`');
				
				texture.wrapS = THREE.RepeatWrapping;
				texture.wrapT = THREE.RepeatWrapping;
				
				var textureLoader = new THREE.TextureLoader();
				
				var material = new THREE.MeshPhongMaterial({
					map: texture,
					`+Normal+`
					`+Specular+`
					`+AmbientOcclusion+`
				});
				
				var loader = new THREE.ColladaLoader( loadingManager );
				loader.load( '`+Mesh+`', function ( collada ) {

					collada.scene.traverse(function (node) {
						if (node.isMesh) {
							node.material.map = material.map;
						}
					});
				
					models.push(collada.scene);
					

				} );
	
				`+PropLoaderString+`

				//

				var ambientLight = new THREE.AmbientLight( 0xcccccc, 0.4 );
				scene.add( ambientLight );

				var directionalLight = new THREE.DirectionalLight( 0xffffff, 0.8 );
				directionalLight.position.set( 1, 1, 0 ).normalize();
				scene.add( directionalLight );

				//

				renderer = new THREE.WebGLRenderer();
				renderer.setPixelRatio( window.devicePixelRatio );
				renderer.setSize( 400, 400 );
				container.appendChild( renderer.domElement );

				window.addEventListener( 'resize', onWindowResize, false );

			}

			function onWindowResize() {

				camera.aspect = 1;
				camera.updateProjectionMatrix();

				renderer.setSize( 400, 400 );

			}

			function animate() {

				requestAnimationFrame( animate );

				render();

			}

			function render() {
				onWindowResize()
			
				var delta = clock.getDelta();

				for (var i = 0; i < models.length; i++) {
					models[i].rotation.z += delta * 0.5;
				}

				renderer.render( scene, camera );

			}

		</script>
	`)
}
