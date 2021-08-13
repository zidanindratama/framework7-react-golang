import React, { Component, useEffect, useState } from "react";
import WithListLoading from "../components/withListLoading";
import axios from "axios";
import { List, ListItem } from "framework7-react";

const endpoint = `http://localhost:8080/mahasiswa`;

const Percobaan = () => {
  const [appState, setAppState] = useState("");

  useEffect(() => {
    fetch(endpoint)
      .then((res) => res.json())
      .then((mahasiswa) => {
        setAppState({ mahasiswa: mahasiswa });
      });
  }, [setAppState]);
  console.log(appState);

  return (
    <>
      <List>
        {appState &&
          appState.mahasiswa.data.map((mhs) => {
            return (
              <ListItem
                key={mhs.nim}
                link={`/about/${mhs.nama}/${mhs.nim}`}
                title={mhs.nama}
                nama={mhs.nama}
                nim={mhs.nim}
              />
            );
          })}
      </List>
    </>
  );
};

export default Percobaan;
