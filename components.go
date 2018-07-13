package main

var ComponentHelp = map[string]string{
	"Decay/Active": "If false, the entity will not do any decaying",
	"Decay/SinkingAnim": "If true, the entity will decay in a ship-like manner",
	"Decay/DelayTime": "Time to wait before starting to sink, in seconds",
	"Decay/SinkRate": "Initial rate of sinking, in metres per second",
	"Decay/SinkAccel": "Acceleration rate of sinking, in metres per second per second",
	
	"Footprint": "Approximation of the entity's shape, for collision detection and outline rendering. Shapes are flat horizontal squares or circles, extended vertically to a given height.",
	"Footprint/MaxSpawnDistance": "Farthest distance units can spawn away from the edge of this entity",
	"Footprint/Height": "Vertical extent of the footprint (in metres)",
	"Footprint/Square": "Set the footprint to a square of the given size",
	"Footprint/Square.width": "Size of the footprint along the left/right direction (in metres)",
	"Footprint/Square.height": "Size of the footprint along the front/back direction (in metres)",
	"Footprint/Circle": "Set the footprint to a circle of the given size",
	"Footprint/Circle.radius": "Radius of the footprint (in metres)",

	"Obstruction": "Causes this entity to obstruct the motion of other units.",
	"Obstruction/Active" : "If false, this entity will be ignored in collision tests by other units but can still perform its own collision tests",
	"Obstruction/BlockMovement" : "Whether units should be allowed to walk through this entity",
	"Obstruction/BlockPathfinding" : "Whether the long-distance pathfinder should avoid paths through this entity. This should only be set for large stationary obstructions",
	"Obstruction/BlockFoundation" : "Whether players should be unable to place building foundations on top of this entity. If true, BlockConstruction should be true too",
	"Obstruction/BlockConstruction" : "Whether players should be unable to begin constructing buildings placed on top of this entity",
	"Obstruction/DeleteUponConstruction" : "Whether this entity should be deleted when construction on a buildings placed on top of this entity is started.",
	"Obstruction/DisableBlockMovement" : "If true, BlockMovement will be overridden and treated as false. (This is a special case to handle foundations)",
	"Obstruction/DisableBlockPathfinding" : "If true, BlockPathfinding will be overridden and treated as false. (This is a special case to handle foundations)",
	"Obstruction/ControlPersist" : "If present, the control group of this entity will be given to entities that are colliding with it.",
	
	"Ownership": "Allows this entity to be owned by players.",
	
	"Position": "Allows this entity to exist at a location (and orientation) in the world, and defines some details of the positioning.",
	"Position/Anchor": "Automatic rotation to follow the slope of terrain",
	"Position/Altitude": "Height above terrain in metres",
	"Position/Floating": "Whether the entity floats on water",
	"Position/FloatDepth": "The depth at which an entity floats on water (needs Floating to be true)",
	"Position/TurnRate": "Maximum graphical rotation speed around Y axis, in radians per second",
	
	"RallyPointRenderer": "Displays a rally point marker where created units will gather when spawned",
	"RallyPointRenderer/MarkerTemplate": "Template name for the rally point marker entity (typically a waypoint flag actor)",
	"RallyPointRenderer/LineTexture": "Texture file to use for the rally point line",
	"RallyPointRenderer/LineTextureMask": "Texture mask to indicate where overlay colors are to be applied (see LineColor and LineDashColor)",
	"RallyPointRenderer/LineThickness": "Thickness of the marker line connecting the entity to the rally point marker",
	"RallyPointRenderer/LinePassabilityClass": "The pathfinder passability class to use for computing the rally point marker line path",
	
	"Selectable": "Allows this entity to be selected by the player.",
	"Selectable/EditorOnly": "If this element is present, the entity is only selectable in Atlas",
	"Selectable/Overlay": "Specifies the type of overlay to be displayed when this entity is selected",
	"Selectable/Overlay/AlwaysVisible": "If this element is present, the selection overlay will always be visible (with transparency and desaturation)",
	"Selectable/Overlay/Texture": "Displays a texture underneath the entity.",
	"Selectable/Overlay/Texture/MainTexture": "Texture to display underneath the entity. Filepath relative to art/textures/selection/.",
	"Selectable/Overlay/Texture/MainTextureMask": "Mask texture that controls where to apply player color. Filepath relative to art/textures/selection/.",
	"Selectable/Overlay/Outline": "Traces the outline of the entity with a line texture.",
	"Selectable/Overlay/Outline/LineTexture": "Texture to apply to the line. Filepath relative to art/textures/selection/.",
	"Selectable/Overlay/Outline/LineTextureMask": "Texture that controls where to apply player color. Filepath relative to art/textures/selection/.",
	"Selectable/Overlay/Outline/LineThickness": "Thickness of the line, in world units.",
	
	"UnitMotion": "Provides the unit with the ability to move around the world by itself.",
	"UnitMotion/WalkSpeed": "Basic movement speed (in metres per second)",
	"UnitMotion/PassabilityClass": "Identifies the terrain passability class (values are defined in special/pathfinder.xml)",
	
	"VisualActor": "Display the unit using the engine's actor system.",
	"VisualActor/Actor": "Filename of the actor to be used for this unit",
	"VisualActor/FoundationActor": "Filename of the actor to be used the foundation while this unit is being constructed",
	"VisualActor/Foundation": "Used internally; if present, the unit will be rendered as a foundation",
	"VisualActor/ConstructionPreview": "If present, the unit should have a construction preview",
	"VisualActor/DisableShadows": "Used internally; if present, shadows will be disabled",
	"VisualActor/ActorOnly": "Used internally; if present, the unit will only be rendered if the user has high enough graphical settings.",
	"VisualActor/SelectionShape/Bounds": "Determines the selection box based on the model bounds",
	"VisualActor/SelectionShape/Footprint": "Determines the selection box based on the entity Footprint component",
	"VisualActor/SelectionShape/Box": "Sets the selection shape to a box of specified dimensions",
	"VisualActor/SelectionShape/Cylinder": "Sets the selection shape to a cylinder of specified dimensions",
}

const ComponentPaletteRaw = `<Entity>

<Decay>
	<Active>false</Active>
	<SinkingAnim>false</SinkingAnim>
	<DelayTime>0</DelayTime>
	<SinkRate>1</SinkRate>
	<SinkAccel>0</SinkRate>
</Decay>

<Footprint>
	<Square width='3.0' height='3.0'/>
	<Height>0.0</Height>
	<MaxSpawnDistance>8</MaxSpawnDistance>
</Footprint>

<Footprint>
	<Circle radius='0.5'/>
	<Height>0.0</Height>
	<MaxSpawnDistance>8</MaxSpawnDistance>
</Footprint>

<Minimap>
	<Type>food, wood, stone, metal, structure, unit or support</Type>
	<Color r="255" g="255" b="255"/> 
</Minimap>

<Obstruction>
	<Static width="1.5" height="1.5"/>
	<Active>true</Active>
	<BlockMovement>true</BlockMovement>
	<BlockPathfinding>false</BlockPathfinding>
	<BlockFoundation>true</BlockFoundation>
	<BlockConstruction>true</BlockConstruction>
	<DeleteUponConstruction>false</DeleteUponConstruction>
	<DisableBlockMovement>false</DisableBlockMovement>
	<DisableBlockPathfinding>false</DisableBlockPathfinding>
	<ControlPersist/>
</Obstruction>

<Obstruction>
	<Unit/>
	<BlockMovement>true</BlockMovement>
	<BlockPathfinding>false</BlockPathfinding>
	<BlockFoundation>true</BlockFoundation>
	<BlockConstruction>true</BlockConstruction>
	<DeleteUponConstruction>false</DeleteUponConstruction>
	<DisableBlockMovement>false</DisableBlockMovement>
	<DisableBlockPathfinding>false</DisableBlockPathfinding>
	<ControlPersist/>
</Obstruction>

<Obstruction>
	<Obstructions>
		<anyName x="0" y="0" width="1.5" height="1.5"/>
	</Obstructions>
	<BlockMovement>true</BlockMovement>
	<BlockPathfinding>false</BlockPathfinding>
	<BlockFoundation>true</BlockFoundation>
	<BlockConstruction>true</BlockConstruction>
	<DeleteUponConstruction>false</DeleteUponConstruction>
	<DisableBlockMovement>false</DisableBlockMovement>
	<DisableBlockPathfinding>false</DisableBlockPathfinding>
	<ControlPersist/>
</Obstruction>

<OverlayRenderer/>
<Ownership/>

<Position>
	<Anchor>upright, pitch, roll or pitch-roll</Anchor>
	<Altitude>0.0</Altitude>
	<Floating>false</Floating>
	<FloatDepth>0.0</FloatDepth>
	<TurnRate>6.0</TurnRate>
</Position>

<RallyPointRenderer>
	<MarkerTemplate>special/rallypoint</MarkerTemplate>
	<LineThickness>0.75</LineThickness>
	<LineStartCap>flat, round, sharp or square</LineStartCap>
	<LineEndCap>flat, round, sharp or square</LineEndCap>
	<LineDashColor r='255' g='255' b='255'></LineDashColor>
	<LinePassabilityClass>default</LinePassabilityClass>
	
	<LineTexture>/path/to/texture</LineTexture>
	<LineTextureMask>/path/to/mask</LineTextureMask>
</RallyPointRenderer>

<RangeOverlayManager/>
<RangeOverlayRenderer/>

<Selectable>
	<Overlay>
		<Texture>
			<MainTexture>/path/to/texture</MainTexture>
			<MainTextureMask>/path/to/mask</MainTextureMask>
		</Texture>
		<Outline>
			<LineTexture>/path/to/texture</LineTexture>
			<LineTextureMask>/path/to/mask</LineTextureMask>
			<LineThickness>1</LineThickness>
		</Outline>
		<AlwaysVisible disable=""/>
	</Overlay>
	<EditorOnly disable=""/>
</Selectable>

<TerritoryInfluence>
	<Root>false</Root>
	<Weight>65535</Weight>
	<Radius>1</Radius>
</TerritoryInfluence>

<UnitMotion>
	<WalkSpeed>1.0</WalkSpeed>
	<PassabilityClass>default</PassabilityClass>
	<FormationController>false</FormationController>
	<Run>
		<Speed>1.0</Speed>
		<Range>1.0</Range>
		<RangeMin>0.0</RangeMin>
		<RegenTime>1.0</RegenTime>
		<DecayTime>0.1</DecayTime>
	</Run>
</UnitMotion>

<Vision>
	<Range>1</Range>
	<RevealShore>false</RevealShore>
</Vision>

<VisualActor>
	<Actor>/path/to/actor</Actor>
	<FoundationActor>/path/to/foundation</FoundationActor>
	<Foundation disable=""/>
	<ConstructionPreview disable=""/>
	<DisableShadows disable=""/>
	<ActorOnly disable=""/>
	<SilhouetteDisplay>false</SilhouetteDisplay>
	<SilhouetteOccluder>false</SilhouetteOccluder>
	<VisibleInAtlasOnly>false</VisibleInAtlasOnly>
	
	<SelectionShape>
		<Bounds/>
		<Box width="1.0" height="1.0" depth="1.0"/>
		<Cylinder radius="1.0" height="1.0">
	</SelectionShape>
</VisualActor>

<Identity>
	<Civ>Put the civilisation name here</Civ>
	<SpecificName>A more specific name</SpecificName>
	<Icon>path/to/icon</Icon>
</Identity>

</Entity>`
