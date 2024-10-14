public class Cono extends Solido {
//OVERVIEW: classe che modella un oggetto di tipo Cono
    private double raggio;
    private double altezza;

//CONSTRUCTOR
    public Cono(String tipo, double raggio, double altezza){
    //MODIFIES: this
    //EFFECTS: istanzia un oggetto di tipo sfera    
        super(tipo);
        this.raggio = raggio;
        this.altezza = altezza;
    }

//METHODS


    @Override
    public int compareTo(Solido o) {
        double v = this.volume();
        double ov = o.volume();
        
        if(v > ov)
            return -1;
        else if(v > ov) 
            return 1;
        else
            return 0;
    }

    @Override
    public double volume() {
        double tmp = Math.PI * Math.pow(this.raggio, 2) * this.altezza;
        return tmp / 3;
    }

    
}
