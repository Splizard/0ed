package main

import "github.com/icza/gowut/gwu"

type Widgets struct {
	gwu.Window
}

func MoreWidgets(w gwu.Window) Widgets {
	return Widgets{Window:w}
}

//Disable the back button of the browser.
func (w Widgets) DisableBackButton() {
	w.Window.Add(gwu.NewHTML(`
	<script>
    
		(function (global) { 

			if(typeof (global) === "undefined") {
				throw new Error("window is undefined");
			}

			var _hash = "!";
			var noBackPlease = function () {
				global.location.href += "#";

				// making sure we have the fruit available for juice (^__^)
				global.setTimeout(function () {
					global.location.href += "!";
				}, 50);
			};

			global.onhashchange = function () {
				if (global.location.hash !== _hash) {
					global.location.hash = _hash;
				}
			};

			global.onload = function () {            
				noBackPlease();

				// disables backspace on page except on input fields and textarea..
				document.body.onkeydown = function (e) {
					var elm = e.target.nodeName.toLowerCase();
					if (e.which === 8 && (elm !== 'input' && elm  !== 'textarea')) {
						e.preventDefault();
					}
					// stopping event bubbling up the DOM tree..
					e.stopPropagation();
				};          
			}

		})(window);
    
    </script>
	`))
}

func NewSoundPlayer(url string) gwu.HTML {
	return gwu.NewHTML(`
	
	<audio controls preload="metadata" style=" width:300px;">
	<source src="`+url+`" type="audio/ogg">
	</audio><br />
	
	`)
}
