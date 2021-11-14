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

package hashx_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "github.com/speps/go-hashids"
    "github.com/w6d-io/hashx"
)

var _ = Describe("Check hash id behaviour", func() {
    Context("", func() {
        BeforeEach(func() {
        })
        AfterEach(func() {
        })
        It("get a hash from id", func() {
            length := 16
            secret := "forty-two"
            hashx.SetHash(&secret, &length)
            h, err := hashx.HashId2String(6)
            Expect(err).ShouldNot(HaveOccurred())
            Expect(h).To(Equal("kyRev14jV4gZPqwo"))
        })
        It("fails due to field missing", func() {
            hashx.SetHash(nil, nil)
            _, err := hashx.HashId2String(42)
            Expect(err).To(HaveOccurred())
            Expect(err.Error()).To(Equal("request field missing - salt and minLength must be set by calling SetHash"))
        })
        It("fails due to shortness alphabet", func() {
            secret := "secret"
            length := 3
            hashx.SetAlphabet("abcdef")
            hashx.SetHash(&secret, &length)
            _, err := hashx.HashId2String(42)
            Expect(err).To(HaveOccurred())
            Expect(err.Error()).To(Equal("request field missing - salt and minLength must be set by calling SetHash: alphabet must contain at least 16 characters"))
        })
        It("fails due to negative id", func() {
            length := 16
            secret := "forty-two"
            hashx.SetHash(&secret, &length)
            hashx.SetAlphabet(hashids.DefaultAlphabet)
            _, err := hashx.HashId2String(-6)
            Expect(err).Should(HaveOccurred())
            Expect(err.Error()).To(Equal("negative number not supported"))
        })

    })
})
