package PROVAESAME;



import java.util.ArrayList;
import java.util.Iterator;
import java.util.Scanner;

public class Vetreria implements Iterable<Contenitore>{
//OVERVIEW: classe che modella una Vetreria, una vetreria  un insieme di Contenitori di diverso tipo
    private ArrayList<Contenitore> contenitori;

//CONSTRUCTOR
    public Vetreria(){
        this.contenitori = new ArrayList<Contenitore>();
    }

//METHODS
    public void addContenitore(Contenitore c){
        this.contenitori.add(c);
    }

    public void removeContenitore(Contenitore c){
        this.contenitori.remove(c);
    }
    public void ordina(){
    //MODIFIES: this.contenitori
    //EFFECTS: riordina i contenitori in base alla loro capienza in ordine decrescente
        this.contenitori.sort(null);
    }

    @Override
    public Iterator<Contenitore> iterator() {
    //EFFECTS: ritorna un itaratore sugli elementi della lista contenitori
        return new Iterator<Contenitore>(){
            int current = 0;
            @Override
            public boolean hasNext() {
                return this.current < Vetreria.this.contenitori.size();
            }

            @Override
            public Contenitore next() {
               return Vetreria.this.contenitori.get(current);
            }

            
        };
    }

    public Vetreria getContenitori(String tipo){
    //MODIFIES: this.contenitori
    //EFFECTS: metodo che estrae dalla vetreria tutti i contenitori con il liquido di un certo tipo,
        Vetreria v = new Vetreria();
        Iterator<Contenitore> i = this.contenitori.iterator();
        while(i.hasNext()){
            Contenitore tmp = i.next();
            if(tmp.getNome().equals(tipo)){
                v.addContenitore(tmp);
                i.remove();
            }
        }
        return v;
    }

    public void distribuisci() throws IncompatibleLiquidsException{
    //MODIFIES: this.contenitori
    //EFFECTS: 
        this.ordina(); //ordino in ordine decrescente se non è già stato fatto
        for (int i = 0; i < this.contenitori.size()-1; i++) {
            Contenitore c1 = this.contenitori.get(i);
            for (int j = i+1; j < this.contenitori.size(); j++) {
                Contenitore c2 = this.contenitori.get(j);
                try {
                    c1.versaLiquido(c2);
                } catch (IncompatibleLiquidsException e) {
                    throw new IncompatibleLiquidsException(e.getMessage());
                }  
            }
        }

        /* 
        for (Contenitore c1 : this.contenitori) {
            for (Contenitore c2 : this.contenitori) {
                try {
                    c1.versaLiquido(c2);
                } catch (IncompatibleLiquidsException e) {
                    throw new IncompatibleLiquidsException(e.getMessage());
                }
            }
        }
        */
    }


//MAIN
    public static void main(String[] args){
        Vetreria v = new Vetreria();

        Scanner s = new Scanner(System.in);

        while(s.hasNextLine()){
            Contenitore cTmp = null;
            String nomeLiquido = s.next();
            double quantitàLiquido = s.nextDouble();

            String tipo = s.next();

            if(tipo.equals("Cuboide")){
                double a = s.nextDouble();
                double b = s.nextDouble();
                double c = s.nextDouble();

                try {
                    cTmp = new Cuboide(quantitàLiquido, nomeLiquido, a, b, c);
                } catch (ExceededCapacityException e) {
                    System.out.println("Il cuboide ha capienza: " + cTmp.getVolume() + e.getMessage());
                }
            }else if(tipo.equals("Sfera")){
                double raggio = s.nextDouble();
                try {
                    cTmp = new Sfera(quantitàLiquido, nomeLiquido, raggio);
                } catch (ExceededCapacityException e) {
                    System.out.println("La sfera ha capienza: " + cTmp.getVolume() + e.getMessage());
                }
            }else{
                double raggio = s.nextDouble();
                double altezza = s.nextDouble();

                try {
                    cTmp = new Cilindro(quantitàLiquido, nomeLiquido, raggio, altezza);
                } catch (ExceededCapacityException e) {
                    System.out.println("Il cilindro ha capienza: " + cTmp.getVolume() + e.getMessage());
                }
            }
            v.contenitori.add(cTmp);
        }
        s.close();

        //System.out.println(v.contenitori.size());

        Iterator  i = v.iterator();
        while(i.hasNext()){
            System.out.println(i.next().toString());
        }
        
        try {
            v.distribuisci();   
        } catch (IncompatibleLiquidsException e) {
            System.out.println(e.getMessage());
        }
    }
}
