package model

import (
    "fmt"
    "testing"
)

func TestAnalysisVisitor_Visit(t *testing.T) {
    //customer := NewEnterpriseCustomer(`eros`)
   // individualCustomer := NewIndividualCustomer(`eros1`)
    c := new(CustomerCol)
    c.Add(&IndividualCustomer{})
    v := ServiceRequestVisitor{}
    c.Accept(&v)
    fmt.Println(c)
}

func TestCustomerCol_Accept(t *testing.T) {

}

func TestCustomerCol_Add(t *testing.T) {

}

func TestEnterpriseCustomer_Accept(t *testing.T) {

}

func TestIndividualCustomer_Accept(t *testing.T) {

}

func TestNewEnterpriseCustomer(t *testing.T) {

}

func TestNewIndividualCustomer(t *testing.T) {

}

func TestServiceRequestVisitor_Visit(t *testing.T) {

}
