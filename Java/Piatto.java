public class Piatto{
//OVERVIEW: classe che rappresenta un piatto di un menu
    private String nome;
    private String tipo;
    private double costo;

//CONSTRUCTORS
    public Piatto(String nome, String tipo, double costo) {
    //MODIFIES: this
    //EFFECTS: istanzia un oggetto di tipo piatto
        this.nome = nome;
        this.tipo = tipo;
        this.costo = costo;
    }

//METHODS
    public String getNome() {
        return nome;
    }
    public void setNome(String nome) {
        this.nome = nome;
    }
    public String getTipo() {
        return tipo;
    }
    public void setTipo(String tipo) {
        this.tipo = tipo;
    }
    public double getCosto() {
        return costo;
    }
    public void setCosto(double costo) {
        this.costo = costo;
    }
    
    @Override
    public String toString() {
        return this.getNome() + " " + this.getTipo() + " "  + this.getCosto()+ "$";
    }
    
}