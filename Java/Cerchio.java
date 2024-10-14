public class Cerchio implements Figura{
//OVERVIEW: classe che modella un cerchio immutabile
    final double raggio;

//CONSTRUCTOR
    public Cerchio(double raggio){
        this.raggio = raggio;
    }

    @Override
    public double perimetro() {
        return (this.raggio * 2) * Math.PI;
    }
    
    @Override
    public String toString() {
        return "Cerchio " + this.raggio;
    }

    @Override
    public double area() {
        // TODO Auto-generated method stub
        return 0;
    }
}
