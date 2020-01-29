import { Injectable, OnInit } from "@angular/core";
import { HttpClient } from "@angular/common/http";
import { BehaviorSubject } from "rxjs";
import { Item } from "../models/item";

@Injectable({
  providedIn: "root"
})
export class ItemService {
  items: Item[] = [];
  itemAdded: boolean = false;
  constructor(private http: HttpClient) {}
  static baseUrl: string = "http://localhost:8080/";

  getAllItems(): Item[] {
    return this.items;
  }

  getItems() {
    this.http
      .get(ItemService.baseUrl + "items")
      .subscribe({ next: this.setItems.bind(this) });
  }

  setItems(data: Item[]) {
    this.items = [...data];
  }

  createItem() {
    this.items = [
      ...this.items,
      { id: undefined, title: undefined, details: undefined }
    ];
  }
}
