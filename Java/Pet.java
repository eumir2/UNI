public abstract class Pet{
//OVERVIEW: classe che modella un animale domestico
    private String nome;

//CONSTRUCTORS
    public Pet(String nome){
    //REQUIRES: nome deve essere != null
    //MODIFIES: this
    //EFFECTS: istanzia un oggetto Pet con il proprio nome
        this.nome = nome;
    }
    
//METHODS
    abstract void verso();
        //EFFECTS: stampa il verso dell'animale

    public String getNome(){
        return this.nome;
    }


}

