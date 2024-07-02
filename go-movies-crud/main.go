package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastName"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for _, item := range movies {

		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}import styles from './antheRsTable.module.css'
import Parser from 'html-react-parser'
import { useEffect, useState } from 'react'
const AntheComplexTable = ({ difficultyanalysis }) => {
    // const [rowData, setRowData] = useState([])
    // const [tableHeader, setTableHeader] = useState([])
    // const [subOccurrence, setSubOccurrence] = useState()
    // let count = 0
    const [complexlevel, setScore] = useState([])
    useEffect(() => {
        let arr = difficultyanalysis
        setScore(arr)
    }, [difficultyanalysis])

    console.log('complexity', complexlevel)
    // useEffect(() => {
    //     let head = []
    //     Object.keys(props.data.dla_sorted[0]).map((item) => {
    //         head.push(item)
    //     })
    //     setTableHeader(head)

    //     let score = []
    //     let subjectOccurrence = []
    //     props.data.dla_sorted.map((item) => {
    //         let tableData = []
    //         head.map((ele) => {
    //             tableData.push(item[ele])
    //         })
    //         subjectOccurrence.push(item.Subject)
    //         score.push(tableData)
    //     })
    //     let obj = {}
    //     let mp = new Map()
    //     subjectOccurrence.map((item) => {
    //         if (mp.has(item)) mp.set(item, mp.get(item) + 1)
    //         else mp.set(item, 1)
    //     })
    //     for (const [key, value] of mp.entries()) {
    //         obj[key] = value
    //     }
    //     setSubOccurrence(obj)
    //     setRowData(score)
    // }, [])

    return (
        <div>
        {difficultyanalysis?.length > 0 ? (
        <div className={styles.tableWrapperComplex}>
          
                <div style={{ minWidth: '400px' }}>
                    <div className={styles.difficut_analysis_section}>
                        <h5>COMPLEXITY Level Analysis</h5>
                        <p>
                            Based on Comparison of Your Marks with Highest Scorer Marks
                            and Average Scorer Marks
                        </p>
                    </div>
                    <table className={styles.complexityTable}>
                        <thead className={styles.antheCThead}>
                            <tr className={styles.antheCTr}>
                            <th>Type of Questions</th>
                                 {difficultyanalysis && difficultyanalysis[0]?.subject.map((item) => {
                                return (
                                    <th
                                        className={styles.cth}
                                        key={item}
                                    >
                                        {item}
                                    </th>
                                )
                            })} 
                            
                            </tr>
                        </thead>
                         <tbody className={styles.antheCTbody}>
                        {/* {rowData.map((item, id) => {
                            if (id > count) {
                                count += subOccurrence[item[0]]
                            }
                            return (
                                <tr
                                    key={'rowData' + id}
                                    className={styles.tableRowData}
                                >
                                    {id == count ? (
                                        <td
                                            rowSpan={subOccurrence[item[0]]}
                                            className={styles.complexTableSubject}
                                        >
                                            {[item[0]]}
                                        </td>
                                    ) : null}
                                    {item.map((ele, id) => {
                                        return (
                                            <td
                                                key={'complexCol' + id}
                                                style={
                                                    id == 0
                                                        ? { display: 'none' }
                                                        : { padding: '8px 3px' }
                                                }
                                                className={styles.colValue}
                                            >
                                                {ele}
                                            </td>
                                        )
                                    })}
                                </tr>
                            )
                        })} */}
                    </tbody> 
                    </table>
                </div>
           
        </div>
         ) : null}  
          </div>
    )
}

export default AntheComplexTable

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)

		}

	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "4233127", Title: "DDLG", Director: &Director{FirstName: "Dhananjay", LastName: "Kapoor"}})
	movies = append(movies, Movie{ID: "2", Isbn: "6546572", Title: "JWM", Director: &Director{FirstName: "Amit", LastName: "Durpa"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting Server at Port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
