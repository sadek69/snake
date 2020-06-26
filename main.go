// uveli4ivat ee pri poedaniii slu4ainih to4ek
// pods4et koli4estva sjedenih to4ek i vivod etogo libo srazy libo posle porazenija
// konec igri pro popadanii na samo sebja liniei to4ek ,zvuki?!
package main

import (
  "fmt"
  "time"
  "math/rand"
  "github.com/tanema/amore"
  "github.com/tanema/amore/audio"
  "github.com/tanema/amore/gfx"
  "github.com/tanema/amore/keyboard"
)
type Points struct {
    x, y []float32
}

var (
 speed float32 = 1.0
 snakex []float32
 snakey []float32
 width   float32=1024
 height  float32=768
 vector float32 = 0
 snake Points
 status float32 = 0
 oldvector float32 = 0
 apple Points
 eat, _ = audio.NewSource("audio/lazer.wav", true)
 dead, _ = audio.NewSource("audio/bomb.wav", true)
)

func main() {
  rand.Seed(time.Now().UnixNano())
  //zadaem pervie startovie koordinati
  snake.x=append(snake.x,width/2)
  snake.y=append(snake.y,height/2)
  snake.x=append(snake.x,width/2)
  snake.y=append(snake.y,height/2+1)
  amore.Start(update, draw)
}

func update(deltaTime float32) {
  if keyboard.IsDown(keyboard.KeyEscape) {
		amore.Quit()
	}
  if keyboard.IsDown(keyboard.KeyUp) && oldvector != 3 {
     vector = 1
  }
  if keyboard.IsDown(keyboard.KeyRight) && oldvector != 4 {
     vector = 2
  }
  if keyboard.IsDown(keyboard.KeyDown) && oldvector != 1 {
     vector = 3
  }
  if keyboard.IsDown(keyboard.KeyLeft) && oldvector != 2 {
     vector = 4
  }
  //proverjat ne peresekli li granicu ekrana perekidivaet na druguju storonu
  bordercheck()
  //otve4aet za izmenenii koordinat massiva to4ek tela zmei v napravlenii vector
  updatesnak(vector,status)
  //geneacija jablok
  createapple()
  //proverka nenado li sjest jabloko
  checkeat()
  //proverka konca igri
  checkdead()
}

func draw() {
  snakedraw()
  appledraw()
  snakedead()
}
//funkcija risuet to4ki tela 4ervjaka
func snakedraw() {
  gfx.SetPointSize(4)
  gfx.SetColor(204, 204, 204, 255)
  gfx.Points(snake.x[0],snake.y[0])
  gfx.SetColor(16, 170, 16, 255)
  gfx.SetPointSize(4)
	for n:=1; n < len(snake.x); n++{
    gfx.Points(snake.x[n],snake.y[n])
  }
}
//funkcija katoroja izmenjaet koordinati to4ek massiva po zadanomu vektoru
//nedolzna pozvoljat dvigatsja cervju nazad
func updatesnak(x float32,status float32) {
  if status ==0 {
    //dvizenie v pravo
     if x==2  {
       if len(snake.x) > 1{
         //udaljaem poslednii element
         snake.x = snake.x[0:len(snake.x)-1]
         snake.y = snake.y[0:len(snake.y)-1]
         //dobavljaem novii element v napravlenii vector to est vperedi
         snake.x = append([]float32{snake.x[0]+speed}, snake.x...)
         snake.y = append([]float32{snake.y[0]}, snake.y...)
       } else {
          for n:=0; n < len(snake.x); n++ {
             snake.x[n]+=speed
          }
        }
       oldvector=2
     }
     //dvizenie v vniz
     if x==3 {
       if len(snake.x) > 1 {
       //udaljaem poslednii element
       snake.x = snake.x[0:len(snake.x)-1]
       snake.y = snake.y[0:len(snake.y)-1]
       //dobavljaem novii element v napravlenii vector to est vperedi
       snake.x = append([]float32{snake.x[0]}, snake.x...)
       snake.y = append([]float32{snake.y[0]+speed}, snake.y...)
       } else {
         for n:=0; n < len(snake.y); n++ {
           snake.y[n]+=speed
       }
     }
       oldvector=3
     }
     //dvizenie v levo
     if x==4 {
       if len(snake.x) > 1 {
          //udaljaem poslednii element
          snake.x = snake.x[0:len(snake.x)-1]
          snake.y = snake.y[0:len(snake.y)-1]
          //dobavljaem novii element v napravlenii vector to est vperedi
          snake.x = append([]float32{snake.x[0]-speed}, snake.x...)
          snake.y = append([]float32{snake.y[0]}, snake.y...)
     } else {
          for n:=0; n < len(snake.x); n++ {
             snake.x[n]-=speed
          }
      }
      oldvector=4
     }
     //dvizenie vverh
     if x==1 {
       if len(snake.x) > 1 {
       //udaljaem poslednii element
       snake.x = snake.x[0:len(snake.x)-1]
       snake.y = snake.y[0:len(snake.y)-1]
       //dobavljaem novii element v napravlenii vector to est vperedi
       snake.x = append([]float32{snake.x[0]}, snake.x...)
       snake.y = append([]float32{snake.y[0]-speed}, snake.y...)
       } else {
       for n:=0; n < len(snake.y); n++ {
        snake.y[n]-=speed
        }
       }
       oldvector=1
     }
   }
}
//funkcija proverjat ne peresekli li granicu i perebrasivaet na drugoi krai ekrana
func bordercheck() {
  for n:=0; n < len(snake.x);n++{
    if snake.x[n]<0 {
      snake.x[n]=width
    }
    if snake.x[n]>width {
      snake.x[n]=0
    }
    if snake.y[n]<0 {
      snake.y[n]=height
    }
    if snake.y[n]>height{
      snake.y[n]=0
    }
 }
}
///funkcija katoroja generiruet pole slu4ainih jablo4ek
//i avtogeneracija novih jablok posle poedanija vseh
func createapple() {
  if len(apple.x)==0 || len(apple.x)==1 {
    rnd:=rand.Intn(25)
    for n:=0;n<rnd;n++ {
      apple.x=append(apple.x,float32(rand.Intn(1024)))
      apple.y=append(apple.y,float32(rand.Intn(768)))
    }
  }
}

//rusiet jabloki
func appledraw() {
  gfx.SetPointSize(5)
  gfx.SetColor(255, 0, 0, 100)
  gfx.SetPointSize(4)
	for n:=1; n < len(apple.x); n++{
    gfx.Points(apple.x[n],apple.y[n])
  }
}

//funkcija katoroja otve4aet za proverku ne sjel li 4erv jabloko i esli sjel to dobavljaet emu
func checkeat() {
  for n:=0;n<len(apple.x);n++ {
    //tut kostili stob  obleg4it sjedenie !
    if snake.x[0]==apple.x[n] && snake.y[0]==apple.y[n] ||  snake.x[0]-1==apple.x[n] && snake.y[0]-1==apple.y[n] || snake.x[0]+1==apple.x[n] && snake.y[0]+1==apple.y[n] ||  snake.x[0]-1==apple.x[n] && snake.y[0]+1==apple.y[n] || snake.x[0]+1==apple.x[n] && snake.y[0]-1==apple.y[n] {
      //udalenie elementa N iz slice jabloka posle sjedanija
      apple.x = append(apple.x[:n], apple.x[n+1:]...)
      apple.y = append(apple.y[:n], apple.y[n+1:]...)
      //dobavlenie element v slice chervjaka
      //nado opredelit po gorizontali prirost chervja ili po vertikali
      if snake.x[len(snake.x)-1]>snake.x[len(snake.x)-2] {
        snake.x = append(snake.x,snake.x[len(snake.x)-1]+1)
        snake.y = append(snake.y,snake.y[len(snake.y)-1])
        } else {
          snake.x = append(snake.x,snake.x[len(snake.x)-1])
          snake.y = append(snake.y,snake.y[len(snake.y)-1]+1)
        }
        //zvuk sjedenija ljuboi
        eat.Play()
      }
  }
}
//funkcija katoroja proverjaet ne sjel li cherv sam sebja,eto
//kogda golova cervja peresekaetsja s telom chervja
func checkdead() {
  for z:=1;z<len(snake.x);z++ {
    if snake.x[0]==snake.x[z] && snake.y[0]==snake.y[z] {
      status=1
      dead.Play()
    }

  }
}
//risuem konec igri
func snakedead() {
  if status==1 {
    gfx.Print("GAME OVER",50, 50)
    gfx.Print("GAME OVER",140, 140)
    gfx.Print("GAME OVER",160, 160)
    gfx.Print("Snake long :",165, 165)
    gfx.Print(fmt.Sprint(len(snake.x)),170, 170)
  }
}
