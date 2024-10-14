import java.time.LocalDate;
import java.util.Scanner;


public class Main {
    public static void main(String[] args) {
        Scanner s = new Scanner(System.in);

        //ArrayList<Prodotto> prodotti = new ArrayList<Prodotto>();

        System.out.println("Inserisci un prodotto nel formato: `<nome> <costo> alimentare <datascadenza>` oppure `<nome> <costo> nondeperibile riciclabile nonriciclabile`");
        while(s.hasNext()){
            Prodotto p = null;
            String nome = s.next();
            double costo = s.nextDouble();
            String tipo = s.next();

            if(tipo.equals("alimentare")){
                String[] dataTmp = s.next().split("-");
                int giorni = Integer.parseInt(dataTmp[0]);
                int mese = Integer.parseInt(dataTmp[1]);
                int anno = Integer.parseInt(dataTmp[2]);
                LocalDate l = LocalDate.of(anno,mese,giorni);
                
                p = new ProdottoAlimentare(nome, costo, l);
            }else if(tipo.equals("nondeperibile")){
                String riciclo = s.next();
                if(riciclo.equals("riciclabile"))
                    p = new ProdottoNonDeperibile(nome, costo, true);
                else if(riciclo.equals("nonriciclabile"))
                    p = new ProdottoNonDeperibile(nome, costo, false); 
            }
           System.out.println("Prezzo con sconto: " + p.costo()); 
           //System.out.println(p.getCosto());
        }

    }
    
}
