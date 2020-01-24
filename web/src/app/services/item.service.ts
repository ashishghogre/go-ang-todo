import { Injectable, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BehaviorSubject } from 'rxjs';
import { Item } from '../models/item';

@Injectable({
  providedIn: 'root'
})
export class ItemService {

  static items: Item[];
  static itemAdded: boolean = false;
  constructor(private http: HttpClient) {};
  static baseUrl: string = "http://localhost:8080/";

  getItems() {
    this.http.get(ItemService.baseUrl + "items").subscribe({next: this.setItems});
  }
  
  setItems(data: Item[]){
    console.log(data)
    ItemService.items = data;
    console.log(ItemService.items)
  }

  createItem() {
    console.log(this);
    ItemService.items = [...ItemService.items,{id:undefined, title: undefined}]
  }
}
