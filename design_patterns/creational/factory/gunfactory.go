package main

import "fmt"

func getGun(gunType string) (IGun, error) {
    if gunType == "ak47" {
        return newAK47(), nil
    }
    if gunType == "musket" {
        return newShotgun(), nil
    }
    return nil, fmt.Errorf("Wrong gun type passed")
}