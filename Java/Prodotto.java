public class Prodotto {
//OVERVIEW: classe che modella un prodotto generico
    private String nome;
    private double costo;

//CONSTRUCTOR
    public Prodotto(String nome, double costo){
    //MODIFIES: THIS
    //EFFECTS: istanzia un prodotto generico, con nome e costo
        this.nome = nome;
        this.costo = costo;
    }

//METHODS
    public String getNome() {
        return nome;
    }

    public void setNome(String nome) {
        this.nome = nome;
    }

    public double getCosto() {
        return costo;
    }

    public void setCosto(double costo) {
        this.costo = costo;
    }
    

    public double costo(){
    //EFFECTS: restituisce il costo del prodotto generico scontato del 5%
        return this.getCosto() - ((this.getCosto() / 100) * 5);
    }
}
