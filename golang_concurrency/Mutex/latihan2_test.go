package mutex

/*
Latihan 2: Membaca dan Menulis dengan RWMutex
Buatlah program yang menggunakan sync.RWMutex untuk mengizinkan banyak goroutine membaca variabel bersama secara bersamaan, tetapi hanya satu goroutine yang boleh menulis pada satu waktu.

Petunjuk:

Gunakan RLock() dan RUnlock() untuk operasi baca.
Gunakan Lock() dan Unlock() untuk operasi tulis.
Buat variabel bersama, lalu jalankan beberapa goroutine yang membaca dan beberapa yang menulis ke variabel tersebut.
*/
