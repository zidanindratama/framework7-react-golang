/// <reference types="cypress" />

const initialItems = [
  {
    name: "181910207",
    phone_number: "Ardi",
  },
  {
    name: "181910208",
    phone_number: "Naufal",
  },
];

const getItems = () =>
  cy
    .request("http://localhost:8080/mahasiswa")
    .its("body")
    .then((resp) => {
      return resp.data;
    });

describe("Get All Mahasiswa", () => {
  it("Get All Mahasiswa", () => {
    cy.request("GET", "http://localhost:8080/mahasiswa")
      .its("headers")
      .its("content-type")
      .should("include", "application/json; charset=utf-8");
  });
});

describe("Create A New Mahasiswa", () => {
  it("Create A New Mahasiswa", () => {
    cy.request("POST", "http://localhost:8080/mahasiswa", {
      id: 1,
      nim: "181910305",
      nama: "zidan indratama",
    })
      .its("headers")
      .its("content-type")
      .should("include", "application/json; charset=utf-8");
  });
});

describe("Get A Detail Of Mahasiswa", () => {
  it("Get A Detail Of Mahasiswa", () => {
    getItems().each((item) => {
      cy.request(`http://localhost:8080/mahasiswa/${item.nim}`)
        .its("body")
        .then((value) => {
          console.log(value.data);
          expect(value.data).to.have.all.keys("id", "nim", "nama");
        });
    });
  });
});
