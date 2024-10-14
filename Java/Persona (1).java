public class Persona {
//OVERVIEW: classe che modella una persona generica
    private String nome;

//CONSTRUCTOR
    public Persona(String nome){
    //MODIFIES: this
    //EFFECTS: istanzia una persona con un nome
        this.nome = nome;
    }
//METHODS

    public String getNome() {
        return nome;
    }  
    
}
