public class Dottore extends Persona{
//OVERVIEW: classe che modella un Dottore
    private String specializzazione;
    private double parcella;

//CONSTRUCTOR
    public Dottore(String nome, String specializzazione, double parcella){
    //MODIFIES: this
    //EFFECTS: istanzia un oggetto di tipo Dottore
        super(nome);
        this.specializzazione = specializzazione;
        this.parcella = parcella;
    }

    
//METHODS
    public String getSpecializzazione() {
        return specializzazione;
    }

    public void setSpecializzazione(String specializzazione) {
        this.specializzazione = specializzazione;
    }

    public double getParcella() {
        return parcella;
    }

    public void setParcella(double parcella) {
        this.parcella = parcella;
    }
    
}
