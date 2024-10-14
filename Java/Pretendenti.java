import java.util.ArrayList;
import java.util.Iterator;
import java.util.NoSuchElementException;

public class Pretendenti implements Iterable<Integer>{
    //OVERVIEW:
    ArrayList<Integer> pretendenti;

    

    //CONSTRUCTORS
    public Pretendenti(int n)throws IllegalArgumentException{
    //MODIFIES: tihs
    //EFFECTS: inizializza this cn n pretendenti, lancia IllegalArgumentException se n <= 0 
        if(n <= 0){
            throw new  IllegalArgumentException("Ci devono essere dei pretendenti");
        }

        this.pretendenti = new ArrayList<Integer>();
        for (int i = 1; i <= n; i++) {
            this.pretendenti.add(i);
        }
    }

    @Override
    public Iterator<Integer> iterator(){
    //EFFECTS: restituisce un iterator che genera nuovi pretendenti da cancellare
    //MODIFIES: l'iteratore può modificare this eliminando i pretendenti

        return new Iterator<Integer>(){

            int current = 0; //indice del pretendente restituito
            boolean removed = true;
            boolean direction = true; //true: sx => dx   false: dx => sx
           
            @Override
            public boolean hasNext() {
            //EFFECTS: restituisce true quando ci sono ancora pretendenti eliminabili, false se ne è rimasto uno solo
                return Pretendenti.this.pretendenti.size() > 1;
            }

            @Override
            public Integer next() throws NoSuchElementException{
            //MODIFIES: this
            //EFFECTS: restituisce un nuovo pretendente da cancellare e toglie il flag "rimosso"
            //         se ne è rimasto uno solo lancia NoSuchElementException  
                if(!(hasNext()))
                    throw new NoSuchElementException("Non ci sono pretendenti");   
                
                if(direction){
                    this.current += 2;
                } else{
                    this.current -= 2;
                }

                if(this.current >= Pretendenti.this.pretendenti.size() -1){
                    this.direction = false;
                    this.current = 2 * (Pretendenti.this.pretendenti.size() - 1) - this.current; 
                }

                if(this.current <= 0){
                    this.direction = true;
                    this.current = -this.current;
                }
                this.removed = false;
                return Pretendenti.this.pretendenti.get(this.current);
            }

            @Override
            public void remove() throws IllegalStateException{
            //MODIFIES: Pretendenti.this, this
            //EFFECTS: rimuove il pretendente restituito da next(), aggiorna l'elemento corrente, imposterà il flag "rimosso" 
            //         se non ho già rimosso un elemento o non ho ancora chiamato next(), lancia IllegalStateException
                if(this.removed)
                    throw new IllegalStateException("Non è possibile rimuovere l'elemento");
                
                Pretendenti.this.pretendenti.remove(this.current);
                removed = true;

                if(!this.direction)
                    this.current--;
                
                if(this.current == 0)
                    this.direction = true;
            
                if(this.current == Pretendenti.this.pretendenti.size() - 1)
                    this.direction = false;

                
            }

            @Override
            public String toString() {
                String dirString = direction ? "avanti" : "indietro";

                return "Si conta da " + Pretendenti.this.pretendenti.get(this.current) + " direzione " + dirString;
            }

        };
    }

        
    //METHODS
    public int last() throws IllegalStateException{
    //EFFECTS: restituisce l'ultimo pretendente rimasto
        if(pretendenti.size() != 1){
            throw new IllegalStateException("Deve esserci un solo pretendente per chiamare last");
        }   
        
        return this.pretendenti.get(0);
    }
    
    @Override
    public String toString(){
        String ret = "Pretendenti: ";
        for (Integer i : this.pretendenti) {
            ret += i + " ";
        }   
        return ret;
    }

    public static void main(String[] args) {
        int n = Integer.parseInt(args[0]);

        Pretendenti p = new Pretendenti(n);
        Iterator<Integer> i = p.iterator();

        while(i.hasNext()) {
            System.out.println(p);
            System.out.println(i);
            System.out.println("Eliminato: " + i.next());
            i.remove();
        } 

        System.out.println("il numero " + p.last() + " è il vincitore");
    }
}
