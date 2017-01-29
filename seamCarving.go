package main
import(
	"log"
	"os"
	"image"
	"strconv"
	"image/color"
	"image/jpeg"
	_"image/png"
	_"image/gif"
	"math"
	"math/rand"
)
func min(a,b int) int{
	if a<b{
		return a
	}else{
		return b
	}
}
func max(a,b int) int{
	if a>b{
		return a
	}else{
		return b
	}
}
type RGB [3]float64
func (a *RGB) set(b color.RGBA64,c float64){
	a[0]=float64(b.R)*c
	a[1]=float64(b.G)*c
	a[2]=float64(b.B)*c
}
func (a *RGB) add(b color.RGBA64,c float64){
	a[0]+=float64(b.R)*c
	a[1]+=float64(b.G)*c
	a[2]+=float64(b.B)*c
}
func (a *RGB) d(b RGB) float64{
	return math.Sqrt(a[0]*a[0]+a[1]*a[1]+a[2]*a[2]+b[0]*b[0]+b[1]*b[1]+b[2]*b[2])
}
func eRnd(a [][]color.RGBA64,b [][]float64){
	for x:=range b{
		for y:=range b[0]{
			b[x][y]=rand.Float64()
		}
	}
}
func eSobel(a [][]color.RGBA64,b [][]float64){
	var X,Y RGB
	w,h:=len(b)-1,len(b[0])-1
	W,H:=w-1,h-1
	X.set(a[1][1],-1.)
	Y=X
	X.add(a[0][1],-2.)
	Y.add(a[1][0],-2.)
	b[0][0]=X.d(Y)
	for y:=H;y>0;y--{
		X.set(a[1][y+1],-1.)
		Y=X
		X.add(a[1][y-1],1.)
		X.add(a[0][y-1],2.)
		X.add(a[0][y+1],-2.)
		Y.add(a[1][y],-1.)
		Y.add(a[1][y-1],-2.)
		b[0][y]=X.d(Y)
	}
	X.set(a[1][H],1.)
	X.add(a[0][H],2.)
	Y.set(a[1][H],-1.)
	Y.add(a[1][h],-2.)
	b[0][h]=X.d(Y)
	for x:=W;x>0;x--{
		X.set(a[x+1][1],-1.)
		Y=X
		X.add(a[x][1],-2.)
		X.add(a[x-1][1],-1.)
		Y.add(a[x-1][1],1.)
		Y.add(a[x-1][0],2.)
		Y.add(a[x+1][0],-2.)
		b[x][0]=X.d(Y)
		for y:=H;y>0;y--{
			X.set(a[x-1][y-1],1.)
			X.add(a[x+1][y+1],-1.)
			Y=X
			X.add(a[x][y-1],2.)
			X.add(a[x][y+1],-2.)
			X.add(a[x+1][y-1],1.)
			X.add(a[x-1][y+1],-1.)
			Y.add(a[x-1][y],2.)
			Y.add(a[x+1][y],-2.)
			Y.add(a[x-1][y+1],1.)
			Y.add(a[x+1][y-1],-1.)
			b[x][y]=X.d(Y)
		}
		X.set(a[x-1][H],1.)
		Y=X
		X.add(a[x][H],2.)
		X.add(a[x+1][H],1.)
		Y.add(a[x+1][H],-1.)
		Y.add(a[x-1][h],2.)
		Y.add(a[x+1][h],-2.)
		b[x][h]=X.d(Y)
	}
	X.set(a[W][1],-1.)
	X.add(a[w][1],-2.)
	Y.set(a[W][1],1.)
	Y.add(a[W][0],2.)
	b[w][0]=X.d(Y)
	for y:=H;y>0;y--{
		X.set(a[W][y-1],1.)
		Y=X
		X.add(a[W][y+1],-1.)
		X.add(a[w][y-1],2.)
		X.add(a[w][y+1],-2.)
		Y.add(a[W][y],2.)
		Y.add(a[W][y+1],1.)
		b[w][y]=X.d(Y)
	}
	X.set(a[W][H],1.)
	Y=X
	X.add(a[w][H],2.)
	Y.add(a[W][h],2.)
	b[w][h]=X.d(Y)
}
func eScharr(a [][]color.RGBA64,b [][]float64){
	var X,Y RGB
	w,h:=len(b)-1,len(b[0])-1
	W,H:=w-1,h-1
	X.set(a[1][1],-3.)
	Y=X
	X.add(a[0][1],-10.)
	Y.add(a[1][0],-10.)
	b[0][0]=X.d(Y)
	for y:=H;y>0;y--{
		X.set(a[1][y+1],-3.)
		Y=X
		X.add(a[1][y-1],3.)
		X.add(a[0][y-1],10.)
		X.add(a[0][y+1],-10.)
		Y.add(a[1][y],-3.)
		Y.add(a[1][y-1],-10.)
		b[0][y]=X.d(Y)
	}
	X.set(a[1][H],3.)
	X.add(a[0][H],10.)
	Y.set(a[1][H],-3.)
	Y.add(a[1][h],-10.)
	b[0][h]=X.d(Y)
	for x:=W;x>0;x--{
		X.set(a[x+1][1],-3.)
		Y=X
		X.add(a[x][1],-10.)
		X.add(a[x-1][1],-3.)
		Y.add(a[x-1][1],3.)
		Y.add(a[x-1][0],10.)
		Y.add(a[x+1][0],-10.)
		b[x][0]=X.d(Y)
		for y:=H;y>0;y--{
			X.set(a[x-1][y-1],3.)
			X.add(a[x+1][y+1],-3.)
			Y=X
			X.add(a[x][y-1],10.)
			X.add(a[x][y+1],-10.)
			X.add(a[x+1][y-1],3.)
			X.add(a[x-1][y+1],-3.)
			Y.add(a[x-1][y],10.)
			Y.add(a[x+1][y],-10.)
			Y.add(a[x-1][y+1],3.)
			Y.add(a[x+1][y-1],-3.)
			b[x][y]=X.d(Y)
		}
		X.set(a[x-1][H],3.)
		Y=X
		X.add(a[x][H],10.)
		X.add(a[x+1][H],3.)
		Y.add(a[x+1][H],-3.)
		Y.add(a[x-1][h],10.)
		Y.add(a[x+1][h],-10.)
		b[x][h]=X.d(Y)
	}
	X.set(a[W][1],-3.)
	X.add(a[w][1],-10.)
	Y.set(a[W][1],3.)
	Y.add(a[W][0],10.)
	b[w][0]=X.d(Y)
	for y:=H;y>0;y--{
		X.set(a[W][y-1],3.)
		Y=X
		X.add(a[W][y+1],-3.)
		X.add(a[w][y-1],10.)
		X.add(a[w][y+1],-10.)
		Y.add(a[W][y],10.)
		Y.add(a[W][y+1],3.)
		b[w][y]=X.d(Y)
	}
	X.set(a[W][H],3.)
	Y=X
	X.add(a[w][H],10.)
	Y.add(a[W][h],10.)
	b[w][h]=X.d(Y)
}
func main(){
	if len(os.Args)<5{
		log.Fatal("usage: seamCaving <in.jpg/png/gif> <out.jpg> <width> <height> [curvature]")
	}
	f,e:=os.Open(os.Args[1])
	if e!=nil{
		log.Fatal(e)
	}
	defer f.Close()
	I,_,e:=image.Decode(f)
	if e!=nil{
		log.Fatal(e)
	}
	B:=I.Bounds()
	b:=make([][]float64,B.Max.X)
	a:=make([][]color.RGBA64,B.Max.X)
	for x:=range b{
		b[x]=make([]float64,B.Max.Y)
		a[x]=make([]color.RGBA64,B.Max.Y)
		for y:=range b[0]{
			A,B,C,D:=I.At(x,y).RGBA()
			a[x][y]=color.RGBA64{uint16(A),uint16(B),uint16(C),uint16(D)}
		}
	}
	l:=1
	if len(os.Args)>5{
		l,e=strconv.Atoi(os.Args[5])
		if e!=nil{
			log.Fatal(e)
		}
		if l<1{
			log.Fatal("curvature smaller than 1")
		}
	}
	B.Max.X,e=strconv.Atoi(os.Args[4])
	if e!=nil{
		log.Fatal(e)
	}
	if B.Max.X>len(b[0]){
		log.Fatal("width greater than input")
	}
	B.Max.Y=0
g:	for len(a[0])>B.Max.X{
		eSobel(a,b)
		for x:=len(b)-1;x>0;x--{
			z:=999999
			for y:=len(b[0])-1;y>=0;y--{
				if z>y+l{
					z=max(0,y-l)
					for t:=z+1;t<min(len(b[0])-1,y+l);t++{
						if b[x][t]<b[x][z]{
							z=t
						}
					}
					b[x-1][y]+=math.Min(b[x][min(len(b[0])-1,y+l)],b[x][z])
				}else{
					if y>=l&&b[x][y-l]<b[x][z]{
						z=y-l
					}
					b[x-1][y]+=b[x][z]
				}
			}
		}
		z:=0
		for y:=len(b[0])-1;y>0;y--{
			if b[0][y]>b[0][z]{
				z=y
			}
		}
		b[0]=append(b[0][:z],b[0][z+1:]...)
		a[0]=append(a[0][:z],a[0][z+1:]...)
		for x:=1;x<len(b);x++{
			t:=max(0,z-l)
			for y:=t+1;y<=min(len(b[0])-1,z+l);y++{
				if b[x][y]<b[x][t]{
					t=y
				}
			}
			z=t
			b[x]=append(b[x][:z],b[x][z+1:]...)
			a[x]=append(a[x][:z],a[x][z+1:]...)
		}

	}
	if B.Max.Y<1{
		b=make([][]float64,B.Max.X)
		A:=make([][]color.RGBA64,B.Max.X)
		for x:=range b{
			b[x]=make([]float64,len(a))
			A[x]=make([]color.RGBA64,len(a))
			for y:=range b[0]{
				A[x][y]=a[y][x]
			}
		}
		a=A
		B.Max.Y=B.Max.X
		B.Max.X,e=strconv.Atoi(os.Args[3])
		if e!=nil{
			log.Fatal(e)
		}
		if B.Max.X>len(b[0]){
			log.Fatal("height greater than input")
		}
		goto g
	}
	f,e=os.Create(os.Args[2])
	defer f.Close()
	if e!=nil{
		log.Fatal(e)
	}
	i:=image.NewRGBA64(B)
	for x:=0;x<B.Max.X;x++{
		for y:=0;y<B.Max.Y;y++{
			i.SetRGBA64(x,y,a[y][x])
		}
	}
	jpeg.Encode(f,i,nil)
}

