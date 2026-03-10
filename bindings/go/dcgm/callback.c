/*
 * Copyright IBM Corp. 2018, 2020
 * SPDX-License-Identifier: Apache-2.0
 */

int violationNotify(void* p) {
    int ViolationRegistration(void*);
    return ViolationRegistration(p);
}
