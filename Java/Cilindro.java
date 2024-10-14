package PROVAESAME;


public class Cilindro extends Contenitore{
//OVERVIEW: classe che modella un Cilindro mutabile
    private double raggio;
    private double altezza;

//CONSTRUCTORS
    public Cilindro(double liquido, String nome, double raggio, double altezza) throws ExceededCapacityException {
    //REQUIRES: raggio e altezza devono essere valori maggiori di 0;
    //MODIFIES: this
    //EFFECTS: istanzia un Cilindro con altezza h e raggio
        super(liquido, nome);
        this.raggio = raggio;
        this.altezza = altezza;
        super.setVolume(calcolaVolume());
    }

//METHODS
    @Override
    public double calcolaVolume() {
        return this.altezza * Math.PI * (this.raggio * this.raggio);
    }

    @Override
    public String toString() {
        String tmp = "Cilindro - altezza: " + this.altezza + ", raggio: " + this.raggio + "\n";
        return tmp + super.toString();
    }
    
}
