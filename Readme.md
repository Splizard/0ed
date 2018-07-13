# 0editor
An entity editor for 0ad, modify units, buildings and more!

### Installation
You will need Go installed and then you can run.

```sh
		go get github.com/Splizard/0ed
```

Run 0ed to launch the editor (Make sure your $GOPATH/bin is inside of your $PATH).
Binary releases will come later.

### Notes

This editor is in heavy development and is lacking a lot of features.
Please report bugs to the issue tracker.

### Roadmap

* Actor Editor in the sidebar.
* Adding components and properties to an entity.
* Sound Playback.
* Drag and drop to edit images/meshes/sounds.
* "Ignore parent" option for each component.
* Multi-tab editing.

### Known Issues

* Has not been tested on Windows at all.
* Props are not shown in the correct location.
* Some models refuse to load, others have incorrect textures. WebGL limitations...
* User interface is ugly, don't worry though, I will be working with a designer to improve it.
* The 0ad data path is currently inflexible.
