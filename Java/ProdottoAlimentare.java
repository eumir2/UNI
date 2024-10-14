import java.time.LocalDate;

public class ProdottoAlimentare extends Prodotto{
//OVERVIEW: classe che modella un prodotto alimentare
    private LocalDate dataScadenza;    

//CONSTRUCTORS
    public ProdottoAlimentare(String nome, double costo, LocalDate dataScadenza){
    //MODIFIES: this
    //EFFECTS: istanzia un prodotto alimentare con nome, costo e data di scadenza. 
        super(nome,costo);
        this.dataScadenza = dataScadenza;
    }


//METHODS
    @Override
    public double costo() {
        //System.out.println(this.checkdata());
        if(checkdata())
            return this.getCosto() - ((this.getCosto() / 100) * 30);
        else
            return super.costo();
    }

    public LocalDate getDataScadenza(){
        return this.dataScadenza;
    }

    public boolean checkdata(){
    //EFFECTS: true se dataScadenza inferiore a 10 giorni, false altrimenti
        int value = (LocalDate.now().plusDays(10).compareTo(this.getDataScadenza()));
        //System.out.println(value);
        if(value <= -1)
            return false;
        else
            return true;
    }
}
