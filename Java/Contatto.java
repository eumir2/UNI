public class Contatto {
//OVERVIEW: classe che modella l'oggetto Contatto, che rappresenta un elemento di una rubrica e contiene il nome, il cognome, il numero di telefono (opzionale) e l'indirizzo mail (opzionale) della persona.
    private String nome;
    private String cognome;
    private String numero;
    private String mail;

    
//CONSTRUCTORS
    public Contatto(String nome, String cognome, String numero, String mail) {
    //MODIFIES: this
    //EFFECTS: crea un nuovo oggetto di tipo contatto; 
    //REQUIRES: tutti i parametri devono essere compatibili   
        this.nome = nome;
        this.cognome = cognome;
        this.numero = numero;
        this.mail = mail;
    }

//METHODS
    public String getNome() {
        return nome;
    }
    public void setNome(String nome) {
        this.nome = nome;
    }
    public String getCognome() {
        return cognome;
    }
    public void setCognome(String cognome) {
        this.cognome = cognome;
    }
    public String getNumero() {
        return numero;
    }
    public void setNumero(String numero) {
        this.numero = numero;
    }
    public String getMail() {
        return mail;
    }
    public void setMail(String mail) {
        this.mail = mail;
    }

    @Override
    public boolean equals(Object obj) {

        if(obj == null)
            return false;
        
        if(!(obj instanceof Contatto))
            return false;

        Contatto c = (Contatto)obj;

        if (!(this.nome.equals(c.nome)))
            return false;

        if(!(this.cognome.equals(c.cognome)))
            return false;

        if(!(this.numero.equals(c.numero)))
            return false;
        
        if(!(this.mail.equals(c.mail)))
            return false;

        return true;
    }

    @Override
    public String toString() {
        return "Contatto [nome=" + nome + ", cognome=" + cognome + ", numero=" + numero + ", mail=" + mail + "]";
    }

    

    

}
