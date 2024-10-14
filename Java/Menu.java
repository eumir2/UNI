import java.util.ArrayList;
import java.util.Iterator;
import java.util.Scanner;

public class Menu implements Iterable {
//OVERVIEW: classe che modella una serie di oggetti Piatto appartenenti a un menu

    private ArrayList<Piatto> piatti;

//CONSTRUCTORS
    public Menu(){
    //MODIFIES: this
    //EFFECTS: istanzia un menu vuoto
        this.piatti = new ArrayList<Piatto>();
    }

//METHODS
    public void addPiatto(Piatto p) throws IllegalArgumentException{
    //MODIFIES: this.piatti
    //EFFECTS: aggiunge un piatto al menu, se già presente lancia IllegalArgumentException
        if (this.piatti.contains(p))
            throw new IllegalArgumentException("Piatto già presente");
        else    
            this.piatti.add(p);
    }
    
    public void removePiatto(Piatto p){
    //MODIFIES: this.piatti
    //EFFECTS: elimina un piatto al menu, se non presente lancia IllegalArgumentException
        if(this.piatti.contains(p))
            this.piatti.remove(p);
        else
            throw new IllegalArgumentException("Piatto non presente");
    }

   /*  public String toString() {
        String tmp = "";
        for (Piatto piatto : piatti) {
            tmp += piatto.toString() + "\n";
        }
        return tmp;
    }*/

    @Override
    public Iterator iterator() {
    //EFFECTS: genera un iterator che restituisce un piatto del menu

        return new Iterator<Piatto>(){
            int current = 0;

            @Override
            public boolean hasNext() {
            //EFFECTS: restituisce true se ci sono ancora dei piatti nel menu, e false se non ce ne sono
                if(this.current < Menu.this.piatti.size())
                    return true;
                else
                    return false;
            }

            @Override
            public Piatto next() {
                Piatto p = Menu.this.piatti.get(this.current);
                this.current ++;
                return p;
            }
            
        };
    }

    public Iterator<Piatto> iterator(String s) {
    //EFFECTS: genera un iterator che restituisce un piatto del menu
    
            return new Iterator<Piatto>(){
                int current = 0;
    
                @Override
                public boolean hasNext() {
                //EFFECTS: restituisce true se ci sono ancora dei piatti nel menu, e false se nonc ce ne sono
                    if(this.current < Menu.this.piatti.size())
                        return true;
                    else
                        return false;
                }
    
                
                public Piatto next() {
                    Piatto p = Menu.this.piatti.get(this.current);
                    this.current ++;
                    if(p.getTipo().equals(s))
                        return p;
                    return null;
                }    
            };        
        }

//MAIN
    public static void main(String[] args) {
        Scanner s = new Scanner(System.in);
        Menu m = new Menu();

        System.out.println("Aggiungi i piatti nel formato: nome tipo costo (terminare la lettura con CTRL+D):");
        while(s.hasNext()){
            String nome = s.next();
            String tipo = s.next();
            String tmp = s.next();
            double costo = Double.parseDouble(tmp.substring(0,tmp.length()-1));
            Piatto p = new Piatto(nome, tipo, costo);
            m.addPiatto(p);
        }

        Iterator<Piatto> i = m.iterator("secondo");
        System.out.println("Secondi:");
        while(i.hasNext()){
            Piatto p = i.next();
            if(p != null)
                System.out.println(p.toString());
        }

        System.out.println();
        System.out.println("-------------------");
        System.out.println();

        Iterator<Piatto> j = m.iterator("primo");
        System.out.println("Primi:");
        while(j.hasNext()){
            Piatto p = j.next();
            if(p != null)
                System.out.println(p.toString());
        }
    }


}
