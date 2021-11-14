/*
Copyright 2020 WILDCARD SA.

Licensed under the WILDCARD SA License, Version 1.0 (the "License");
WILDCARD SA is register in french corporation.
You may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.w6d.io/licenses/LICENSE-1.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is prohibited.
Created on 14/11/2021
*/

package hashx

import (

    "github.com/speps/go-hashids"
    "github.com/w6d-io/x/errorx"
)

var (
    salt *string
    minLength *int
    alphabet  = hashids.DefaultAlphabet
)

func SetHash(secret *string, length *int) {
    salt = secret
    minLength = length
}

func SetAlphabet(alpha string) {
    alphabet = alpha
}

func HashId2String(id int) (string, error) {
    if salt == nil || minLength == nil {
        return "", &errorx.Error{
            Code:       "crypto_hash_request_field_missing",
            Message:    "request field missing - salt and minLength must be set by calling SetHash",
        }
    }
    hd := hashids.NewData()
    hd.Salt = *salt
    hd.MinLength = *minLength
    hd.Alphabet = alphabet
    h, err := hashids.NewWithData(hd)
    if err != nil {
        return "", &errorx.Error{
            Cause: err,
            Code:       "crypto_hash_request_field_missing",
            Message:    "request field missing - salt and minLength must be set by calling SetHash",
        }
    }
    secret, err := h.Encode([]int{id})
    if err != nil {
        return "", err
    }
    return secret, nil
}
