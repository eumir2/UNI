import java.time.LocalDate;

public class ProdottoNonDeperibile extends Prodotto{
//OVERVIEW: classe che modella in Prodotto non deperibile
    private boolean flag;

//CONSTRUCTOR
    public ProdottoNonDeperibile(String nome, double costo, boolean flag){
    //MODIFIES: this
    //EFFECTS: istanzia un prodotto con nome, costo, data di scadenza e 
    //         flag che indica se il prodotto è riciclabile o meno
        super(nome, costo);
        this.flag = flag;
    }


    @Override
    public double costo() {
        if(flag)
            return this.getCosto() - ((this.getCosto() / 100 ) * 10);
        else
            return super.costo();
    }
    
}
