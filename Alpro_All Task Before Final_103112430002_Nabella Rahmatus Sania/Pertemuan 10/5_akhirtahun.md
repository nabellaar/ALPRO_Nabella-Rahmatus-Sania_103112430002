Program Menghitung Diskon dan Cashback

Kamus:
- total : integer
- buatKartu : boolean

Algoritma:
1. Input nilai total dan buatKartu
2. Jika total >= 100000 maka:
     a. total ← total - (total / 10)
3. Jika total >= 200000 dan buatKartu = true maka:
     a. total ← total - 75000
4. Tampilkan "Kartu?" dan nilai buatKartu
5. Tampilkan "Diskon?" dan kondisi total >= 100000
6. Tampilkan "Cashback?" dan kondisi buatKartu = true dan total + 75000 >= 200000
7. Tampilkan "Rp." dan nilai total

End Program