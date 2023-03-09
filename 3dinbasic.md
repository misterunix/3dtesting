## The QB Times

Issue #2

### 3D in QB

Hi folks! It's me, your all-nutcase, 100% freaked out programmer, BlackBird. Nightwolf asked me to write an article about 3D techniques for his online magazine, the QB Times, and that's what you're reading right now. I originally shoud have written this for issue #1 but, as always, I was late. Nevertheless, it's here now, and it will be better then ever! (er... that shouldn't be too hard for a first time I guess)
Since I have been working on a 3D modeller (or editor, or drawing program, whatever), called 3DES Designer (or 3DES for short) for quite some time now, I figured it would be nice to combine the 3D tutorial with a 3DES tutorial, also because I have since received many emails regaring the use of the models created in 3DES. Soon I will release a 3DES SDK (or Software Development Kit) for use with your own programs.

As you might have noticed, the 3DES zipfile came with a small program called LOAD3DP.BAS which explained the 3DP file format that 3DES uses to store it's models. However, this is not sufficent for most users and therefore I will try to explain a little about the 3DES models in this article.

But first you'll have to understand the basics of 3D graphics, so here we go... Firstly, three dimensional objects consist of points in a 3D space, and for the totally braindead among us, three dimensional points have three coordinates: the X, Y and Z coordinates.

In maths classes at school (at least mine) they teach you Z is the horizontal axis and Y is the axis that points into the depth... bummer for them, 'cause I hereby treat Z as the 'depth' axis, and X and Y as the 2D axis that lie on your computer screen... let's use a nice ASCII drawing to support that:
(just imagine that the drawing is 3D and that the Z-axis points into the depth)

```
 o--------> x-axis
 |\
 | \
 |  \ z-axis
 |  
 \/  y-axis
```

Mind that 3D coordinates are notated in the following order: X,Y,Z
Also, the arrows in the drawing above point to the positive values relative to the center of the cooridinate system (being o in the drawing), eg: negative Z values lie in front of the center, positive Z values behind it.
Got that? good. Now how do you display those points? Well, since your screen is only 2D, you can't just plot a pixel at, say 10,4,12... (or PSET (10,4,12),15) you could try... but it definitively won't work.
So what you need to do is 'convert' the 3D coordinates to 2D ones that fit onto your screen (this process is called translation).
But how? It's quite simple, just follow this formula:
```
FlatX = X * 256 / Z
FlatY = Y * 256 / Z
```        
Where FlatX and FlatY are the 2D coordinates that you can plot to the screen right away. Noticed the number 256 in there? Right, now that is the value that controls one of the most important things in the world of 3D graphics, perspective (meaning things get smaller the further they are away from the viewer). The larger the value, the less perspective in your object, keep that in mind. 256 is a pretty good 'default' value. You don't need to change it unless your model's perspective looks exaggerated or something.
Okay, so we've gotten quite far now. You now ought to know how to draw a 3D object consisting of points on a 2D screen...
if that is still unclear, let me clarify it for you, here's a little source code to get you started:
```
-------------8<-------------8<-------------8<-------------8<-------------

'// Ajust this value to load larger models or save memory
CONST numberOfPoints = 100
'// Ajust this value to zoom in or out on the object
CONST zoomLevel = 0

TYPE ThreeDeePoint
	X AS SINGLE
	Y AS SINGLE
	Z AS SINGLE
END TYPE

'// Set up an array of 3D points
DIM MyArray(1 TO numberOfPoints) AS ThreeDeePoint

'// Load a model into the array


'// Draw all the points
FOR pointIdx = 1 TO numberOfPoints
	'// Do not draw points that are behind the screen
	IF MyArray(pointIdx).Z + zoomLevel > 0 THEN
		'// Translate points
		flatX = MyArray(pointIdx).X * 256 / (MyArray(pointIdx).Z + zoomLevel)
		flatY = MyArray(pointIdx).Y * 256 / (MyArray(pointIdx).Z + zoomLevel)
		'// Plot a white dot at the point's translated coordinates
		PSET (flatX, flatX), 15
	END IF
NEXT

-------------8<-------------8<-------------8<-------------8<-------------
```        
And voila! your model is here! Offcourse it still looks like shit, since you only used dots to draw the points.
Most 3D games and applications, like 3DES for example, use face-based models. A face consist of 3 or more points, and when linked together, they form a polygon. Here's another one of my beautiful ASCII drawings to clarify that:
``` 
              o - point 1
             / \
            /   \
 point 4 - o     \
            \     \
   point 3 - o-----o - point 2
```        
The drawing above is a face consisting of four points, to draw it, just translate all four points to 2D, draw a line from point 1 to 2, 2 to 3, 3 to 4 and from 4 back to 1, and hey presto! A face!
Offcourse it's still a wireframe representation, not a solid face, but that's yet another chapter, drawing a filled polygon, and I'm not going to explain that here.
So, with a little creativity, you can apply the face-technique to the sourcecode above, making it a so-called wireframe 3D engine. Hurray!
Still, we are missing the main part of a true 3D engine... manipulation of the model. It's very nice and all, a still of a 3D model, but it's much more fun when the thing actually rotates around it's axis.
3D rotation is basically not very different from 2D rotation, not to say exactly the same.
The only thing you need to do is rotate each point three times, one time for each axis (X,Y and Z).
Rotating a 2D point can be accomplished by using this formula:
```
newX = COS(degrees) * X - SIN(degrees) * Y
newY = SIN(degrees) * X + COS(degrees) * Y
```        
To translate a 3D point, you'll have to treat the rotation around each axis (3 in total) as a 2D rotation using the remaining two coordinates.
For example, to rotate around the X-axis, treat the Y,Z coordinates as if they were 2D coordinates in 2D space.
Assuming you put the above rotation function in a SUBroutine defined like this: SUB rotatePoint (X, Y, newX, newY, degrees), the correct code to rotate a point around the axis would be:
```
-------------8<-------------8<-------------8<-------------8<-------------

'// around X-axis
CALL rotatePoint (Z, Y, newZ, newY, degrees)

'// around Y-axis
CALL rotatePoint (X, Z, newX, newZ, degrees)

'// around Z-axis
CALL rotatePoint (Y, X, newY, newX, degrees)

-------------8<-------------8<-------------8<-------------8<-------------
```        
Then all you have to do is just rotate all points like in the example above, then translate them and then draw the bloody thing =)
Well folks, that is the end of part 1, this series will be continued in future issues of the QB Times, see yers!

This article was written by BlackBird
























## SECTION 1 PART A SUBPART 3 |  3D Programming

### 3D Graphics in BASIC - Part I.1: Introduction 

HEY YOU !

Interrested in programming vector graphics ? Bored of painting 2 dimensional
diagrams. Feared of using C/C++ ? Or are you a math dummie ? 

No panic.

Spend some time (and of course even more time at your PC) with me and join 
me into the wonderful world of ...

**3D Graphics in BASIC (yep, it's possible !!!)**


Before I fill your brain with lot of stuff like vectors, matrix, filled 
polygons or shading techniques in the next five parts (this text included), 
I'll introduce me to my person:

I'm Christian Garms, a german chemistry student and programmer. Last year 
I've made some nice money with Visual Basic programs in MS Excel and that's 
why I think BASIC is the opposite of a dead programming language. I'm a
registered PB user but the examples that will be post in this article should
work with both QB and PB (eventually with minor modifications).

For questions, REMarks or comments send an e-mail to:

<garms@chemie.uni-hamburg.de>


I assume that you're not a BASIC beginner. Hope that you're at least an 
advanced programmer because this article is not a bedtime story. 

But the sweetest fruits are hidden and hard to get.

Most programs will do the trick without any assembly additives to speed up 
the code. If x86-ASM is necessary I'll write it as INLINE Assembler.  
Sorry QB Users but I have not much time to spend for converting INLINE code 
into suitable OBJ-Files.

### 3D Graphics in BASIC - Part I.2: some basic math

The hardest part to understand of 3D graphics is the mathematical background 
that is - politely spoken - abstract.  
But lets begin with an easier entry point. `Point` - that is the right object
to start with. In our three dimensional world all points consist of three
components, the `x-`, `y-` and `z-coordinate`. With these values every point is
strongly determined in his position. But to whom ?  

That is the next thing to be dealt with: **Coordinate systems**!  

For the beginning we start with only one coordinate system, the world 
coordinates. That means all coordinates are related to an absolut center 
somewhere in our real world. E.g.: a value of `x = 0` , `y = 0` and `z = 0` define 
a point exactly in the center of our world. 

For a mathematician every point in the whole 3D world is a vector. That's why
3D graphics is also called vector graphics.  
If we had a point Z who lies in the center of our world than the definition 
of Z will be:

```
Z = (0 0 0)
```

Any other point P with unequal values to zero of `x-`, `y-`, and `z-coordinates`
would look like:

```
P = (<x> <y> <z>)
```

The letters in the brackets are placeholders for the corresponding values.
OK, lets return from the equation thicket to programming.  
For further use we should define our own TYPE of variable, a vector !

```
' Creating own definition
TYPE vector
  x AS INTEGER
  y AS INTEGER
  z AS INTEGER
END TYPE

' Declare p as vector
DIM p AS vector

' Sets p to the center of the world
p.x = 0
p.y = 0
p.z = 0
```

Listing I-2.1

Listing I-2.1 demonstrates the usage of user-defined types. User-defined types makes your programs more structured and better understandable than Listing I-2.2.  
Especially when you have more than one point!  
As you can see I'm using mostly integer arithmetics. That is a common trick 
to speed up the output of 3D graphics enormously.


```
' Define point in the middle of the universe
px% = 0
py% = 0
pz% = 0
```

Listing I-2.2



### 3D Graphics in BASIC - Part I.3: transformations

Now we have the simpliest object: a point. The next question is: How to convert a point in our 3D world - or mathematically spoken a vector - into  a flat pixel on the screen.  
So here comes the moment to introduce you with a new coordinate system - the  eye coordinate system. I think that you, dear reader, will ask WHY. Well, imagine a scenery from any 3D game you have in mind. In most cases of these games there is a craft that you fly, drive or move and others that will be steered by your computer or someone else. You can look in all directions without steering into this directions. This would be impossible if your eye coordinate system is non-existant. In other words: The eye coordinate system allows watching different from the world coordinate systems.  

And now the strategy to convert a point into a pixel:

1. Transformation of the world coordinates into eye coordinates
2. Transformation of eye coordinates into screen coordinates

But some mathematics first. I hope you've got your machet right by your hand and follow me again into the equation jungle. This time it will be harder than last time.

Mathematicians are sometimes lazy to write complex formulas. In the case of transformation of a vector to a new vector in another coordinate system like the transformation of the world coordinate system into the eye coordinate system they simply write:

```
P_eye = P_world * T
```

|P_eye|eye coordinate vector|
|-|-|
|P_world|world coordinate vector|
|T|Transformation operator|

That means: transformation of coordinates is only a "simple" mathematical operation. But I would not go any further now because I've saved that for Part II.  
This time I'll explain the transformation by an example. Once again you must imagine to sit in a craft in our virtual 3D game. Say you're at Position `x=100`, `y=0`, `z=0` and look to the center of our world. If you've reset your nav computer and set the absolute position `(0 0 0)` to your craft (the eye coordinate system) the center of the world now lies at `x=-100`, `y=0`, `z=0`.  
In Summary: the world coordinates of a point will transformed to eye coordinates via the following equations:

```
x_eye = x_world - eyepos_x
y_eye = y_world - eyepos_y
z_eye = z_world - eyepos_z
```

|x_eye|x-coordinate of the point (eye coordinate system)|
|-|-|
|y_eye|y-coordinate of the point (eye coordinate system)|
|z_eye|z-coordinate of the point (eye coordinate system)|
|x_world|x-coordinate of the point (world coordinate system)|
|y_world|y-coordinate of the point (world coordinate system)|
|z_world|z-coordinate of the point (world coordinate system)|
|eyepos_x|position of the watcher (relative to world center)|
|eyepos_y|position of the watcher (relative to world center)|
|eyepos_z|position of the watcher (relative to world center)|


But we gained also a three-dimensional point. How to convert this one into a pixel? Now the mathematician comes in action. And he won't be lazy any more ! He will tell you something about triangles, pyramids ... and you're stuck complete helpless in the thickest formula thicket you could 
think about. If there would be a chalkboard he would easily write it full just for explanations. Simple, isn't it ?

Instead of molto formulos there is THE golden wisdom of every 3D-Programmer:

> "The screen coordinates could be calculated by dividing the x- and y- position through the depth  (z-coordinate)"

In formula speak:

```
x_screen = x_eye / z_eye
y_screen = y_eye / z_eye
```

with    

|x_screen|x-coordinate of the pixel|
|-|-|
|y_screen|y-coordinate of the pixel|
|x_eye|See above|
|y_eye|See above|
|z_eye|See above|

You gain a dimensionless number that must be fit to screen coordinates and 
to the middle of the screen.  
I assume that the width and the height of the screen are given so the formula
results to:

```
x_screen = (x_eye / z_eye) * width + width / 2
y_screen = (y_eye / z_eye) * height + height / 2
```

Now we've got all parts together to write some real 3D stuff.

```
' Simple 3D Object (Pyramid)

' Type declarations
TYPE vector
  x AS INTEGER
  y AS INTEGER
  z AS INTEGER
END TYPE

TYPE pixel
  x AS INTEGER
  y AS INTEGER
END TYPE

' Variable declaration
DIM p(3) AS vector
DIM eye AS vector
DIM s(3) AS pixel
DIM maxx AS INTEGER     ' width of screen
DIM maxy AS INTEGER     ' height of screen

' Screen resolution 
maxx = 640
maxy = 480 

' Read Object Data
FOR i = 0 TO 3
  READ p(i).x
  READ p(i).y
  READ p(i).z
NEXT i

' Definition of object
DATA 30, 1, 1
DATA  1,30, 1
DATA  1, 1,30
DATA -30,-30,-30

' Set Eye position (change if desired)
eye.x = 0
eye.y = 0
eye.z = 100

' Calculate the eye coordinates
FOR i = 0 TO 3
  p(i).x = p(i).x - eye.x
  p(i).y = p(i).y - eye.y
  p(i).z = p(i).z - eye.z
NEXT i

' Calculate screen coordinates
FOR i = 0 TO 3
  s(i).x = (p(i).x / p(i).z) * maxx + maxx / 2
  s(i).y = (p(i).y / p(i).z) * maxy + maxy / 2
NEXT i

' Draw object
CLS
SCREEN 12
FOR i = 0 TO 5
  READ pt1, pt2
  LINE (s(pt1).x,s(pt1).y)-(s(pt2).x,s(pt2).y)
NEXT i

DATA 0, 1, 0, 2, 0, 3, 1, 2, 1, 3, 2, 3
```

Listing I-3.1

OK, folks. Next time I will introduce you to animated vector graphics and
the calclulation with matrix. Stay tuned and I hope that you enjoy this article.



















## 3D Graphics in BASIC - Part II.1: Epilog /

It sounds strange to begin with an epilog but I have to explain some real
important things about the listing PYRAMID.BAS in the last part:

1. You only have to calculate the edges of a polygon. In the case of the pyramid you only have to calculate FOUR (!!!) points. The rest will do the LINEs. This is very good because it reduces the amount of calculations to a minimum and also, of course, the amount of cpu usage.
2. You must have an exact represantation of your 3D object. In our case of the pyramid this is very simple. There are only four points. The definition of the object is located in the DATAs. You need in general a DATA statement for the points and a DATA statement for the connectivity list.  
The connectivity list will instruct the program to draw the right lines tothe right points.  
You must determine every edge exactly of any given 3d object you want to display. This is very time consuming and you can only create smaller objects with a pencil and a paper sheet. For 'bigger' objects (more points) you need a special editor.



### 3D Graphics in BASIC - Part II.2: 3D Animations
   
In PYRAMID.BAS there is only one single picture of a simple object.

Boring, isn't ?

The real 3D effect will only show up if the object will be animated like rotating around an axis or moving in real time. So I want you to show how to get this 'pyramid' into action.

But tons of theory first ...

To make understanding easier for this relative difficult subject because this part is like "Formula Jones and The Raiders of the lost Arc" and you could be easily get lost somewhere in the Amasinus I'll give you an overview of what to do:

1. 3D animations - and of course, animations in general - need the double buffering technique. That's a common used method of displaying and generating pictures simultaneously on different screen (often done by choosing different memory locations of the displaying screen and the drawing screen).  
If you would display and draw the picture on the same screen (the same memory location), the picture might become flickery. With the double buffering technique - and eventually waiting for the vertical retrace interrupt - the animated graphics looks very smooth.
2. Rotating, Scaling and Moving of 3D points could be done with matrix operations. Because matrix operations aren't a simply matter of fact at all I'll explain it here in this article but limited for our purpose.  
When you've worked through this stuff (it's a very thick formula thicket - have you got your machet right by your hand ?) you will see the advantages of this mathematical technique.




### 3D Graphics in BASIC - Part II.4: 3D Object moving
 

Moving - or also called: translation - of an object is done by changing the coordinates of the object. Let's start with a simple example: a point in the 3D world. Moving the point could be done by:

        1. changing the points coordinates:

        obj.x = obj.x + t.x
        obj.y = obj.y + t.y
        obj.z = obj.z + t.z

        With (obj.x/obj.y/obj.z) = 3D point and (t.x/t.y/t.z) =
        translation vector. The translation vector describes how much a
        point is moved in any direction (x,y,z).

        or (very important !)

        2. changing the viewers coordinates:

        eye.x = eye.x + t.x
        eye.y = eye.y + t.y
        eye.z = eye.z + t.z

        With (eye.x/eye.y/eye.z) = viewers' point

The result will be the same: The point will be moved. That's the same
phenomon as if we watched the sunrise. Not the sun is going up but our
planet earth is rotating around his polar axis. We know that the earth is
moving but it looks like the sun is moving.
The listing II-4.1 is the modified example of the last part - PYRAMID.BAS.
Now it shows some motion. The pyramid is bouncing (in fact the viewpoint
is moving) to the viewer and away from him/her.

**************************************************************************
' ---------------------
'   Moving Pyramid
' based on PYRAMID.BAS
' (C) 1996 by Ch. Garms
' ---------------------


' Compiler Instructions

$CPU 80386
$OPTIMIZE SPEED
$LIB GRAPH ON
$ERROR ALL ON
$COMPILE MEMORY


' Creating new TYPEs

TYPE vector
  x AS INTEGER
  y AS INTEGER
  z AS INTEGER
END TYPE

TYPE pixel
  x AS INTEGER
  y AS INTEGER
END TYPE


' Variable declarations

%MAXPT = 3                      ' max. points
%MAXLN = 5          ' max. lines
DIM s(%MAXPT) AS pixel      ' 2D coordinates of Pyramid
DIM eye AS vector       ' viewpoint
DEFINT a-z


' Initializing screen constants

%MAXPOSX = 639          ' max. X-coordinate of screen
%MAXPOSY = 349          ' max. Y-coordinate of screen
%CENTERX = 320          ' center of screen (X-position)
%CENTERY = 175          ' center of screen (Y-position)


' Initializing Viewpoint

eye.x = 15
eye.y = 15
eye.z = 0


' Calculating the eye coordinates & transformation into screen pixels

SUB vec2pix( objpt AS vector, scrpix AS pixel )
  SHARED eye

  DECR objpt.x, eye.x
  DECR objpt.y, eye.y
  DECR objpt.z, eye.z

  scrpix.x = (objpt.x / objpt.z) * %MAXPOSX + %CENTERX
  scrpix.y = (objpt.y / objpt.z) * %MAXPOSY + %CENTERY
END SUB


' Switch screens:
' implementation for PB's SCREEN

SUB switchscreen
  STATIC display, drawing

  IF display = drawing THEN
    display = 0
    drawing = 1
  ELSE
    SWAP display, drawing
  END IF

  WAIT &H3DA, 8         ' wait for vertical retrace
  SCREEN 9, 0, display, drawing
  CLS
END SUB


' IMPORTANT: from here starts the nonrecycable code

' Main program

DIM pwork AS vector
WHILE NOT INSTAT
  FOR j = 40 TO 200 STEP 2
    switchscreen
    eye.z = j
    RESTORE objectdata
    FOR i = 0 TO %MAXPT
      READ pwork.x, pwork.y, pwork.z
      vec2pix pwork, s(i)
    NEXT i
    RESTORE connectdata
    FOR i = 0 TO %MAXLN
      READ pt1, pt2
      LINE (s(pt1).x,s(pt1).y) - (s(pt2).x,s(pt2).y)
    NEXT i
  NEXT j
  FOR j = 200 TO 40 STEP -2
    switchscreen
    eye.z = j
    RESTORE objectdata
    FOR i = 0 TO %MAXPT
      READ pwork.x, pwork.y, pwork.z
      vec2pix pwork, s(i)
    NEXT i
    RESTORE connectdata
    FOR i = 0 TO %MAXLN
      READ pt1, pt2
      LINE (s(pt1).x,s(pt1).y) - (s(pt2).x,s(pt2).y)
    NEXT i
  NEXT j
WEND

SCREEN 0


' Object Data & Connectivity list

objectdata:
DATA  30,  1,    1
DATA   1, 30,    1
DATA   1,  1,   30
DATA -30, -30, -30

connectdata:
DATA  0, 1, 0, 2, 0, 3, 1, 2, 1, 3, 2, 3
**************************************************************************
Listing II-4.1: MOVINPYR.BAS

The listing II-4.1 has some nice features. It contains code that's
recycable (you mustn't reinvent the wheel !). For our purposes there are
two new SUBs:

        SUB switschscreen:
        This subroutine flips between two pages in the video mode 9
        (640x375x16 colours). This EGA resolution is more than enough
        for simple vector graphics.

        SUB vec2pix:
        Converts a vector (3D point) to a screen pixel. This SUB is
        resolution independant. You have to define only the viewpoint
        (setting eye.x/eye.y/eye.z) and the screen parameters %MAXPOSX,
        %MAXPOSY,%CENTERX,%CENTERY before your first call.


    ---------------------------------------------------
  /                                                  / |
 ----------------------------------------------------  |
 3D Graphics in BASIC - Part II.5: 3D Object rotating /
 ----------------------------------------------------

There isn't much to say about 3D rotating. Only formulas, formulas,
formulas. I think our friend Formula Jones won't be unhappy if we come to
the point right now:

        Rotating around the x-axis (Global coordinate system):
        x' = x*cos(alpha) - y*sin(alpha)
        y' = x*sin(alpha) + y*cos(alpha)
        z' = z

        Rotating around the y-axis (Global coordinate system):
        x' = x*cos(beta) + z*sin(beta)
        y' = y
        z' = -x*sin(beta) + z*cos(beta)

        Rotating around the z-axis (Global coordinate system):
        x' = x'
        y' = y*cos(gamma) - z*sin(gamma)
        z' = y*sin(gamma) + z*cos(gamma)

        With:
        (x/y/z)         = old point
        (x'/y'/z')      = new point
        alpha           = angle to rotate around x-axis clockwise
        beta            = angle to rotate around y-axis clockwise
        gamma           = angle to rotate around z-axis clockwise

I won't explain the origin of these formulas because that will not fit
into this article. If you're interrested you'll find this very complex
stuff in any "higher" math book.
Listing II-5.1 is an example of use. Our well known pyramid is now
rotating around his z- and x-axis. But the basic program can be easily
changed. If you want to rotate to any other axis then you have only to
change the calls. Just experimentate with this program !

**************************************************************************
' ---------------------
'   Rotating Pyramid
' based on PYRAMID.BAS
' (C) 1996 by Ch. Garms
' ---------------------


' Compiler Instructions

$CPU 80386
$OPTIMIZE SPEED
$LIB GRAPH ON
$ERROR ALL OFF
$FLOAT EMULATE


' Creating new TYPEs

TYPE vector
  x AS INTEGER
  y AS INTEGER
  z AS INTEGER
END TYPE

TYPE pixel
  x AS INTEGER
  y AS INTEGER
END TYPE


' Variable declarations

%MAXPT = 3                          ' max. points
%MAXLN = 5              ' max. lines
%FACTOR = 16384
%ANGLE = 3600               ' max. angles for sinus and cosinus
DIM s(%MAXPT) AS pixel
DIM sinus(%ANGLE) AS SHARED INTEGER ' array for sinus table
DIM cosinus(%ANGLE) AS SHARED INTEGER   ' array for cosinus table
DIM eye AS SHARED vector        ' viewpoint
DIM pwork AS vector
deg2rad! = 1800/3.14152695
DEFINT a-z


' Initializing Sinus table

FOR i = 0 TO %ANGLE
  sinus(i)   = CINT( SIN( i/deg2rad!) * %FACTOR )
  cosinus(i) = CINT( COS( i/deg2rad!) * %FACTOR )
NEXT i


' Screen constants

%MAXPOSX = 639              ' max. X-coordinate of screen
%MAXPOSY = 349              ' max. Y-coordinate of screen
%CENTERX = 320              ' center of screen (X-position)
%CENTERY = 175              ' center of screen (Y-position)


' Clipping constants

%LEFT     = 1
%RIGHT    = 2
%UP       = 4
%DOWN     = 8
%TRUE     = -1
%FALSE    = 0


' Initializing Viewpoint

eye.x = 0
eye.y = 0
eye.z = 150


' Rotating Point around X-Axis
'   objpt : vector in world coordinates (!)
'   alpha : angle to rotate around X-Axis (1 means 0.1 deg)

SUB rotatex( objpt AS vector, alpha AS INTEGER )
  SHARED sinus(), cosinus()
  DIM p AS vector

  p.x = (objpt.x * cosinus(alpha) - objpt.y * sinus(alpha)) / %FACTOR
  p.y = (objpt.x * sinus(alpha) + objpt.y * cosinus(alpha)) / %FACTOR

  objpt.x = p.x
  objpt.y = p.y
END SUB


' Rotating Point around Y-Axis
'   objpt : vector in world coordinates (!)
'   beta  : angle to rotate around Y-Axis (1 means 0.1 deg)

SUB rotatey( objpt AS vector, beta AS INTEGER )
  SHARED sinus(), cosinus()
  DIM p AS vector

  p.x = (objpt.x * cosinus(beta) + objpt.z * sinus(beta)) / %FACTOR
  p.z = (objpt.x * -sinus(beta) + objpt.z * cosinus(beta)) / %FACTOR

  objpt.x = p.x
  objpt.z = p.z
END SUB


' Rotating Point around Z-Axis
'   objpt : vector in world coordinates (!)
'   gamma : angle to rotate around Y-Axis (1 means 0.1 deg)

SUB rotatez( objpt AS vector, gamma AS INTEGER )
  SHARED sinus(), cosinus()
  DIM p AS vector

  p.y = (objpt.y * cosinus(gamma) - objpt.z * sinus(gamma)) / %FACTOR
  p.z = (objpt.y * sinus(gamma) + objpt.z * cosinus(gamma)) / %FACTOR

  objpt.y = p.y
  objpt.z = p.z
END SUB


' Calculating the eye coordinates & transformation into screen pixels
'   objpt : vector in world coordinates (!)
'   scrpix: pixel on screen
' The variable eye (TYPE vector) must be defined before calling this sub.

SUB vec2pix( objpt AS vector, scrpix AS pixel )
  SHARED eye

  DECR objpt.x, eye.x
  DECR objpt.y, eye.y
  DECR objpt.z, eye.z

  scrpix.x = (objpt.x / objpt.z) * %MAXPOSX + %CENTERX
  scrpix.y = (objpt.y / objpt.z) * %MAXPOSY + %CENTERY
END SUB


' Switch screens

SUB switchscreen
  STATIC display, drawing

  IF display = drawing THEN
    display = 0
    drawing = 1
  ELSE
    SWAP display, drawing
  END IF

  WAIT &H3DA, 8 ' wait for vertical retrace
  SCREEN 9, 0, display, drawing
  CLS
END SUB


' IMPORTANT: from here starts the nonrecycable code
' Main program

WHILE NOT INSTAT
  FOR j=0 TO %ANGLE STEP 15
    switchscreen
    RESTORE objectdata
    FOR i = 0 TO %MAXPT
      READ pwork.x, pwork.y, pwork.z
      rotatez pwork, j
      rotatey pwork, j
      vec2pix pwork, s(i)
    NEXT i
    RESTORE connectdata
    FOR i = 0 TO %MAXLN
      READ pt1, pt2
      LINE (s(pt1).x,s(pt1).y) - (s(pt2).x,s(pt2).y)
    NEXT i
  NEXT j
WEND

SCREEN 0


' Object Data & Connectivity list

objectdata:
DATA  30,  0,   0
DATA   0, 30,   0
DATA   0,  0,  30
DATA -30,-30, -30

connectdata:
DATA  0, 1, 0, 2, 0, 3, 1, 2, 1, 3, 2, 3
**************************************************************************
Listing II-5.1: ROTPYR.BAS

The listing II-5.1 has a very nice trick: The sinus and cosinus values are
converted to integers by multiplying with a constant factor (the factor
must be less than 32767) and stored in an integer array. That makes the
calculation of the 3D rotating faster than with floating point math. It is
not a great secret because it is a well used technique for vector graphics
since games like Elite on the C-64. Though the calculation aren't very
precise the screen resolution is so small that calculation errors won't
disturb much.

The listing II-5.1 simplifies the rotating. As you can see all formulas
rotate around an axis of the global coordinate system. But the pyramid is
rotating around a point in the center of the pyramid. The program achieves
this by equalising the center of the object and the center of the global
coordinate system. If a chosen scenery is more complex (e.g. two objects
who rotates differemt) then we come to a new coordinate system which I will
now introduce: The Object coordinate system.
That means: All points of a given object will be defined relative to the
center of the object. To display the object in the global cordinate system
(or: world coordinate system) we have only add the translation vector from
the center of the object to the center of the global corrdinate system to
all points of the object.

For example I will take the single point once more for explanation of this
complex subject:
The point is the center of the object. The relative object coordinate will
be (0/0/0) and the translation vector (x/y/z). To display the point into the
world coordinate system we simply add the translation vector to the object
coordinates so the derived global coordinte point is (x/y/z).

In general:

        world.x = obj.x + transl.x
        world.y = obj.y + transl.y
        world.z = obj.z + transl.z

        with:
        (world.x/world.y/world.z) = world coordinates of object point
        (obj.x/obj.y/obj.z) = object coordinates of object point
        (transl.x/transl.y/transl.z) = translation vector of object

That's the same as translating a 3D point in the world coordinate system.
Now we've defined our object within the object coordinate system we only
have to equalise the object center and the world center in our mind. For
the rotations we take the object coordinates not the world coordinates !
Than we can perform the rotations. To display the object we add the
translation vector of the object center to all object points and convert
the points to screen pixels.


    -------------------------------------------------------------------
  /                                                                   / |
 ---------------------------------------------------------------------  |
 3D Graphics in BASIC - Part II.6: Introductions to Matrix calculations /
 ----------------------------------------------------------------------

Matrix operations aren't a mystical thing. You have not to be a math
genius to understand what matrices are:

        "A Matrix is a represantion of a linear equation"

In other words: A Matrix isn't more than an array of values which contains
the suffixes of any linear equation like:

        a1*x + b1*y + c1*z = d1
        a2*x + b2*y + c2*z = d2
        a3*x + b3*y + c3*z = d3

The corresponding matrix looks as follows:

        |a1 b1 c1|   |d1|
        |a2 b2 c2| = |d2|
        |a3 b3 c3|   |d3|

For our purposes we didn't need more to know. As you have seen our 3D
operations are often performed by linear equations. E.g. translation of a
point is performed by adding the translation vector to a point. If we
write down this equation in a matrix form it will look like:

     Matrix1   Matrix2       Matrix3
        |x|   |1 0 0 t.x|   |x + t.x|
        |y|   |0 1 0 t.y| = |y + t.y|
        |z|   |0 0 1 t.z|   |z + t.z|
        |0|   |0 0 0 0  |   |0      |

That means we only multiply the 3D point (Matrix1) with an operator
(Matrix2) to translate the point. It looks like I want to complicate all.
But the advantage of matrix operations is that you can chain many
3D operations like rotation or translation to only one single matrix for
all points of any object. This will reduce the calculations enormous and
speed up 3D graphics dramatically for larger objects.

OK, guys. Next time I will continue you to explain the calculation with
matrix and go further with filled polygon graphics.
Hope to see you again here.




