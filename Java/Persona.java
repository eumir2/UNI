import java.util.ArrayList;
import java.util.Iterator;

public class Persona implements Iterable<Pet>{
//OVERVIEW: classe che modella la classe persona che possiede diversi animali
    
    private String nome;
    private ArrayList<Pet> animali;

//CONSTRUCTORS
    public Persona(String nome){
    //REQUIRES: nome != null
    //MODIFIES: this
    //EFFECTS: istanzia un oggetto del tipo Persona con un arraylist di Pet vuota 
        this.nome = nome;
        this.animali = new ArrayList<Pet>();
    }
//METHODS

    public String getNome(){
        return this.nome;
    }

    public void addAnimale(Pet p)throws IllegalArgumentException{
    //MODIFIES: this
    //EFFECTS: pet == null -> lancia IllegalArgumentException, altrimenti aggiunge animale
        if(p == null)
            throw new IllegalArgumentException("Animale nullo");
        else
            this.animali.add(p);
    }

    public void removeAnimale(Pet p)throws IllegalArgumentException{
    //MODIFIES: this
    //EFFECTS: pet == null -> lancia IllegalArgumentException, altrimenti rimuove animale se esiste
        if(p == null)
            throw new IllegalArgumentException("Animale nullo");
        else   
            this.animali.remove(p);

    }

    @Override
    public Iterator<Pet> iterator() {
        return new Iterator<Pet>(){
            int curr = 0;
        
            @Override
            public boolean hasNext() {
                return curr < Persona.this.animali.size();
            }

            @Override
            public Pet next() {
                Pet p = Persona.this.animali.get(curr);
                curr++; 
                return p;
            }
            
        };
    }
    @Override
    public boolean equals(Object obj) {
        if(obj instanceof Persona){
            Persona p = (Persona) obj;
            if(this.getNome().equals(p.getNome()))
                return true;
        }
        return false;
    }


}
