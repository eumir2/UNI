
public class Sfera extends Solido{
//OVERVIEW: classe che modella un oggetto Sfera
    private double raggio;

//CONSTRUCTOR
    public Sfera(String tipo, double raggio){
    //MODIFIES: this
    //EFFECTS: istanzia un oggetto di tipo Sfera con il proprio raggio
        super(tipo);
        this.raggio = raggio;
    }

    @Override
    public int compareTo(Solido s) {
        var v = this.volume();
        var sv = s.volume();

        if(v < sv)
            return -1;
        else if(v > sv) 
            return 1;
        else
            return 0;
    }



    @Override
    public double volume() {
        double tmp = 4 * Math.PI;
        double r = Math.pow(this.raggio, 3);
        return (tmp * r) / 3;
    }  
}
