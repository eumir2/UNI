import java.util.ArrayList;
import java.util.Scanner;

public class Dipendente extends Persona{
    //OVERVIEW: classe che modella un dipendente
    private double retribuzione;
    private int annoAssunzione;
    private String codiceId;

//CONSTRUCTOR
    public Dipendente(String nome, double retribuzione, int annoAssunzione, String codiceId) {
    //MODIFIES: this
    //EFFECTS: istanzia un oggetto Dipendente    
        super(nome);
        this.retribuzione = retribuzione;
        this.annoAssunzione = annoAssunzione;
        this.codiceId = codiceId;
    }

//METHODS
    public double getRetribuzione() {
        return retribuzione;
    }

    public void setRetribuzione(double retribuzione) {
        this.retribuzione = retribuzione;
    }

    public int getAnnoAssunzione() {
        return annoAssunzione;
    }

    public void setAnnoAssunzione(int annoAssunzione) {
        this.annoAssunzione = annoAssunzione;
    }

    public String getCodiceId() {
        return codiceId;
    }

    public void setCodiceId(String codiceId) {
        this.codiceId = codiceId;
    }   


//MAIN
    public static void main(String[] args) {
        int anno = Integer.parseInt(args[0]);
        Scanner s = new Scanner(System.in);
        ArrayList<Dipendente> persone = new ArrayList<Dipendente>();
        Dipendente p = null;

        while(s.hasNext()){
            String nome = s.next();
            String codice = s.next();
            int annoS = s.nextInt();
            double salario = s.nextDouble();

            p = new Dipendente(nome, salario, annoS, codice);
            persone.add(p);
        }

        for (Dipendente d : persone) {
            if(d.getAnnoAssunzione() < anno){
                System.out.println(d.getNome());
            }
        }
    }
}
