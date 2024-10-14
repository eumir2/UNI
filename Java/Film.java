import java.time.LocalDate;

abstract class Film {
//OVERVIEW: classe che modella un film generico
    private static int offset = 0;
    private int id;
    private String titolo;
    private LocalDate dataNoleggio;
    private String tipo;

//CONSTRUCTOR
    public Film(String titolo, String tipo, LocalDate dN){
    //MODIFIES: this
    //EFFECTS: istanzia un oggetto di tipo Film
        this.id = offset++;
        this.titolo = titolo;
        this.tipo = tipo;
        this.dataNoleggio = dN;
    }

//METHODS
    public double calcolaPenaleRitardo(){
    //EFFECTS: calcola il costo della penale, se il film è stato noleggiato da meno
    //         di una settimana, lancia IllegalArgumentException
        double tmp = (LocalDate.now().getDayOfMonth() - this.dataNoleggio.getDayOfMonth()) - 7;
        if(tmp > 0)
            return tmp * 2;
        else
            throw new IllegalArgumentException("Film nolegiato da meno di una settimana");
    };

    public static int getOffset() {
        return offset;
    }

    public int getId() {
        return id;
    }

    public String getTitolo() {
        return titolo;
    }

    public LocalDate getDataNoleggio() {
        return dataNoleggio;
    }

    public String getTipo() {
        return tipo;
    }


}
