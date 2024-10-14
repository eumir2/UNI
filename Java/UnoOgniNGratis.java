public class UnoOgniNGratis extends PoliticaSconto {
//OVERVIEW: classe che modella uno sconto gni n articoli
    private int n;

//CONSTRUCTOR
    public UnoOgniNGratis(int numeroArticoli, double prezzoArticolo, int n) {
        super(numeroArticoli, prezzoArticolo);
        this.n = n;
    }

//METHODS
    
    @Override
    public double calcolaSconto() {
        int tmp = super.getNumeroArticoli() % this.n;  //quante volte devo applicare lo sconto
        return (super.getPrezzoArticolo()  * super.getNumeroArticoli()) - (super.getNumeroArticoli() * tmp);
    }
    
}
