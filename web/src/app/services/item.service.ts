import { Injectable, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BehaviorSubject } from 'rxjs';
import { Item } from '../models/item';

@Injectable({
  providedIn: 'root'
})
export class ItemService {

  static items: Item[];

  constructor(private http: HttpClient) {};

  getItems() {
    this.http.get("http://localhost:8080/items").subscribe({next: this.setItems});
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
