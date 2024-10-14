public class Rettangolo implements Figura{
//OVERVIEW: classe che modella un triangolo immutabile
    final double base;
    final double altezza;

//CONSTRUCTOR
    public Rettangolo(double base, double altezza){
        this.base = base;
        this.altezza = altezza;
    }

//METHODS
    @Override
    public double perimetro() {
        return (this.base + this.altezza) * 2;
    }

    @Override
    public String toString() {
        return "Rettangolo " + this.base + " " + this.altezza;
    }
    
}
