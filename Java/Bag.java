import java.util.HashMap;
import java.util.Iterator;
import java.util.NoSuchElementException;

public class Bag<E> implements Iterable<E>{
//OVERVIEw: classe che modella un Multiset/Bag che memorizza gli elementi inseriti 
//          e il numero di volte di inserimento per ciascuno 
    private HashMap<E,Integer> elementi;



//CONSTRUCTORS
    public Bag(){
    //MODIFIES: this
    //EFFECTS: istanzia una BAg/Multiset vuoto 
        this.elementi = new HashMap<>();   
    }

//METHODS
    public void insert(E elem) throws IllegalArgumentException{
    //MODIFIES: this
    //EFFECTS: se elem == null, lancia IllegalArgumentException, se != null insert in this
        if(elem == null)
            throw new IllegalArgumentException("Elemento nullo");
        Integer count = this.elementi.putIfAbsent(elem, 1);
        if(count != null){
            count+=1;
            this.elementi.replace(elem, count);
        }
    }
    
    public void remove(E elem) throws IllegalArgumentException{
    //MODIFIES: this
    //EFFECTS: se elem == null, lancia IllegalArgumentException, se != null remove from this
        if(elem == null)
            throw new IllegalArgumentException("Elemento nullo");
        Integer count = this.elementi.get(elem);

        if(count != null){
            if(count == 1)
                this.elementi.remove(elem);
            else{
                this.elementi.replace(elem, count-1);
            }
        }

    }
 
    @Override
    public String toString() {
        String ret = " ";
        for(E i : this){
            ret += i + " ";
        }
        return ret + "  ]";
    }
   
    @Override
    public Iterator<E> iterator() {
    //EFFECTS: restituisce un iterator di E
    
        return new Iterator<E>(){
            E curr = null;
            int counter = 1;
            Iterator<E> keys = elementi.keySet().iterator();
            @Override
            public boolean hasNext() {
            //EFFECTS: restituisce TRUE se ci sono ancora elementi 
            //         ed il loro contatore è inferiore al numero degli elementi
            //         altrimenti false
                if(curr != null && elementi.get(curr) > counter)
                    return true;
                else if(keys.hasNext())
                    return true;
                
                return false;
            }

            @Override
            public E next() throws NoSuchElementException {
            //MODIFIES: contatore degli elementi
            //EFFECTS: restituisce il prossimo elemento, aggiornando il contatore,
            //         se il contatore supera la dispoibilità, passa al successivo
            //         se non ci sono pù elementi lancia NoSuchElementException
                if(!(this.hasNext()))
                    throw new NoSuchElementException("no more elements");
                
                if(curr != null && elementi.get(curr) > counter)
                    counter++;
                else{
                    keys.next();
                    counter = 1;
                }   
                return curr;
            }

        };
    }
    
}
