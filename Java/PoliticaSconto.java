abstract class PoliticaSconto {
//OVERVIEW: classe che modella una politica sconto per un numero di oggetti dello stesso tipo
    private int numeroArticoli;
    private double prezzoArticolo;

//CONSTRUCTOR
    public PoliticaSconto(int numeroArticoli, double prezzoArticolo){
        this.numeroArticoli = numeroArticoli;
        this.prezzoArticolo = prezzoArticolo;
    }


//METHODS
    public int getNumeroArticoli() {
        return numeroArticoli;
    }

    public void setNumeroArticoli(int numeroArticoli) {
        this.numeroArticoli = numeroArticoli;
    }

    public double getPrezzoArticolo() {
        return prezzoArticolo;
    }

    public void setPrezzoArticolo(double prezzoArticolo) {
        this.prezzoArticolo = prezzoArticolo;
    }

    abstract double calcolaSconto();
    
}
