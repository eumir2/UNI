public class Triangolo implements Figura{
//OVERVIEW: classe che modella un triangolo immutabile
    final double lato1;
    final double lato2;
    final double lato3;

//CONSTRUCTOR
    public Triangolo(double lato1, double lato2, double lato3){
    //MODIFIES: this
    //EFFECTS: istanzia un oggetto Triangolo dati 3 lati 
        this.lato1 = lato1;
        this.lato2 = lato2;
        this.lato3 = lato3;
    }

//METHODS
    @Override
    public double perimetro() {
        return lato1 + lato2 + lato3;
    }

    @Override
    public String toString() {
        return "Triangolo " + this.lato1 + " " + this.lato2 + " " + this.lato3;
    }


    @Override
    public double area() {
        double p = perimetro()/2;
        double tmp = p * (p-lato1) * (p - lato2) * (p-lato3);
        return Math.sqrt(tmp);
    }

}
