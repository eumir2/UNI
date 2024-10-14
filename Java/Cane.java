public class Cane extends Pet {
//OVERVIEW: classe che modella l'oggetto cane
    public Cane(String nome){
    //REQUIRES: nome != null
    //MODIFIES: this
    //EFFECTS: istanzia un oggetto di tipo Cane
        super(nome);
    }

    @Override
    public void verso() {
    //MODIFIES: stdout
    //EFFECTS: stampa il verso dell'animale
        System.out.print("BAU");
    }
}
