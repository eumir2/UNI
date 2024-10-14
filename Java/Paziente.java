public class Paziente extends Persona{
//OVERVIEW: classe che modella un paziente
    private String id;


//CONSTRUCTOR    
    public Paziente(String nome, String id) {
        super(nome);
        this.id = id;
    }

//METHODS
    public String getId() {
        return id;
    }


    public void setId(String id) {
        this.id = id;
    }

}
