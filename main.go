package main

import (
    "fmt"
    "math"
    "os"
)

func mfcc(audioPath string) []float64 {
    // Load the audio file.
    file, err := os.Open(audioPath)
    if err != nil {
        fmt.Println(err)
        return nil
    }
    defer file.Close()

    // Calculate the MFCCs.
    mfccs := calculateMFCCs(file)

    // Return the MFCCs.
    return mfccs
}

func calculateMFCCs(file *os.File) []float64 {
    // Create a short-time Fourier transform (STFT) of the audio file.
    stft := spectrogram(file)

    // Create a Mel filterbank.
    melFilterbank := melFilterbank()

    // Calculate the MFCCs from the STFT and the Mel filterbank.
    mfccs := calculateMFCCsFromSTFT(stft, melFilterbank)

    // Return the MFCCs.
    return mfccs
}

func spectrogram(file *os.File) []float64 {
    // Calculate the Fourier transform of the audio file.
    fft := fourierTransform(file)

    // Create a spectrogram from the Fourier transform.
    spectrogram := make([]float64, len(fft))
    for i := range spectrogram {
        spectrogram[i] = math.Log10(math.Abs(fft[i]))
    }

    // Return the spectrogram.
    return spectrogram
}

func melFilterbank() []float64 {
    // Create a Mel scale.
    melScale := make([]float64, 128)
    for i := range melScale {
        melScale[i] = 1125 * math.log10(1 + (float64(i) / 700))
    }

    // Create a Mel filterbank.
    melFilterbank := make([]float64, len(melScale))
    for i := range melFilterbank {
        melFilterbank[i] = 0
    }

    for i := range melScale {
        for j := 0; j < len(melFilterbank); j++ {
            if melScale[i] >= melFilterbank[j] && melScale[i] < melFilterbank[j+1] {
                melFilterbank[j] += 1
            }
        }
    }

    // Return the Mel filterbank.
    return melFilterbank
}

func calculateMFCCsFromSTFT(stft []float64, melFilterbank []float64) []float64 {
    // Calculate the Mel-frequency cepstral coefficients (MFCCs) from the STFT and the Mel filterbank.
    mfccs := make([]float64, 12)
    for i := range mfccs {
        mfccs[i] = 0
    }

    for i := range mfccs {
        for j := 0; j < len(stft); j++ {
            mfccs[i] += stft[j] * melFilterbank[i]
        }
    }

    // Return the MFCCs.
    return mfccs
}

func main() {
    // Calculate the MFCCs for the audio file at `audioPath`.
    mfccs := mfcc("audio.wav")

    // Print the MFCCs.
    for i := range mfccs {
        fmt.Println(mfccs[i])
    }
}
