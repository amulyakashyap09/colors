# Colors
> A small golang utility to print colorful output on terminal

## **_Installation_**
```
import color "github.com/amulyakashyap09/colors"
```

## **_Usage_**

```
var message string = "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit..."

#bgColor, fgColor, message
color.Custom("black", "yellow", message)

```

## **_Documentation_**

`color.Custom(bgColor, fgColor, message)`   |   `color.Custom` will print the `message` in colors provided as input

`color.Error(message)` | `color.Error` will print the `error` in red & bold

`color.Panic(message)` | `color.Panic` will print the `error` in red & bold & will **exit** the process

`color.Success(message)` | `color.Success` will print the `message` in green

`color.Info(message)` | `color.Info` will print the `message` in blue

`color.Warn(message)` | `color.Warn` will print the `message` in yellow

`color.Yellow(message)` | `color.Yellow` will print the `message` in yellow

`color.Blue(message)` | `color.Blue` will print the `message` in blue

`color.Green(message)` | `color.Green` will print the `message` in green

`color.Red(message)` | `color.Red` will print the `message` in red

`color.Magenta(message)` | `color.Magenta` will print the `message` in magenta

`color.Cyan(message)` | `color.Cyan` will print the `message` in cyan

`color.White(message)` | `color.White` will print the `message` in white

`color.Black(message)` | `color.Black` will print the `message` in black

## **_Author_**
```
Amulya Kashyap
amulyakasyap09@gmail.com
9559974779
```
