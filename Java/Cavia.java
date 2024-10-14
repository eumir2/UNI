public class Cavia extends Pet {
    //OVERVIEW: classe che modella l'oggetto cavia
        public Cavia(String nome){
        //REQUIRES: nome != null
        //MODIFIES: this
        //EFFECTS: istanzia un oggetto di tipo Cavia
            super(nome);
        }
    
        @Override
        public void verso() {
        //MODIFIES: stdout
        //EFFECTS: stampa il verso dell'animale
            System.out.print("SQUIT");
        }
    }
