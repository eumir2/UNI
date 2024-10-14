public class ScontoQuantità{
//OVERVIEW: classe che modella uno sconto sulla quantità
    private int minimo;
    private int percentuale;
    private PoliticaSconto p;


    public ScontoQuantità(int minimo, int percentuale) {
        this.minimo = minimo;
        this.percentuale = percentuale;
    }

    @Override
    public double calcolaSconto() {
        double tmp = super.getNumeroArticoli() * super.getPrezzoArticolo();
        if(super.getNumeroArticoli() > minimo){
            return tmp - (tmp / 100) * this.percentuale;
        }
        return tmp;
    }    
}
