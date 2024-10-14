public class Gatto extends Pet {
    //OVERVIEW: classe che modella l'oggetto gatto

        public Gatto(String nome){
        //REQUIRES: nome != null
        //MODIFIES: this
        //EFFECTS: istanzia un oggetto di tipo Gatto
            super(nome);
        }
    
        @Override
        public void verso() {
        //MODIFIES: stdout
        //EFFECTS: stampa il verso dell'animale
            System.out.print("MIAO");
        }
    }
