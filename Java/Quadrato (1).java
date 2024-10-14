public class Quadrato implements Figura{
//OVERVIEW: classe che modella un quadrato immutabile
    final double lato;

//CONSTRUCTOR
    public Quadrato(double lato){
        this.lato = lato;
    }


//METHODS
    @Override
    public double perimetro() {
        return this.lato * 4;
    }

    @Override
    public String toString() {
        return "Quadrato " + this.lato;
    }

    @Override
    public double area() {
        return this.lato * this.lato;
    }
}
