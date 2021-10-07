<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateKhenthuongkyluatTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('khenthuongkyluat', function (Blueprint $table) {
            $table->id();
            $table->string('name');
            $table->string('hinhthuc');
            $table->tinyInteger('mucdo');
            $table->date('ngayquyetdinh');
            $table->text('ghichu');
            $table->unsignedBigInteger('id_nhanvien');
            $table->foreign('id_nhanvien')->references('id')->on('nhanvien');
            $table->timestamps();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('khenthuongkyluat');
    }
}
