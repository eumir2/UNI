package PROVAESAME;

public class Cuboide extends Contenitore{
//OVERVIEW: classe che modella un cuboide mutabile
    private double a;
    private double b;
    private double c;

//CONSTRUCTORS
    public Cuboide(double liquido, String nome, double a, double b, double c) throws ExceededCapacityException{
    //REQUIRES: a, b e c devono essere valori maggiori di 0;
    //MODIFIES: this
    //EFFECTS: istanzia un cuboide con lati a, b e c;
        super(liquido, nome);
        this.a = a;
        this.b = b;
        this.c = c;
        super.setVolume(calcolaVolume());
    }

//METHODS
    @Override 
    public double calcolaVolume() {
        return this. a * this.b * this.c;
    }

    public String toString() {
        String tmp = "Cuboide -  lati: ";
        tmp += this.a + ", ";
        tmp += this.b + ", ";
        tmp += this.c + "\n";
        return tmp + super.toString();
    }
}
